---
title: "Go Concurrency Patterns: Pipelines and cancellation"
date: 2014-03-13
by:
- Sameer Ajmani
tags:
- concurrency
- pipelines
- cancellation
summary: How to use Go's concurrency to build data-processing pipelines.
---

## Introduction

Go's concurrency primitives make it easy to construct streaming data pipelines
that make efficient use of I/O and multiple CPUs.  This article presents
examples of such pipelines, highlights subtleties that arise when operations
fail, and introduces techniques for dealing with failures cleanly.

## What is a pipeline?

There's no formal definition of a pipeline in Go; it's just one of many kinds of
concurrent programs.  Informally, a pipeline is a series of _stages_ connected
by channels, where each stage is a group of goroutines running the same
function.  In each stage, the goroutines

  - receive values from _upstream_ via _inbound_ channels
  - perform some function on that data, usually producing new values
  - send values _downstream_ via _outbound_ channels

Each stage has any number of inbound and outbound channels, except the
first and last stages, which have only outbound or inbound channels,
respectively.  The first stage is sometimes called the _source_ or
_producer_; the last stage, the _sink_ or _consumer_.

We'll begin with a simple example pipeline to explain the ideas and techniques.
Later, we'll present a more realistic example.

## Squaring numbers

Consider a pipeline with three stages.

The first stage, `gen`, is a function that converts a list of integers to a
channel that emits the integers in the list.  The `gen` function starts a
goroutine that sends the integers on the channel and closes the channel when all
the values have been sent:

{{code "pipelines/square.go" `/func gen/` `/^}/`}}

The second stage, `sq`, receives integers from a channel and returns a
channel that emits the square of each received integer.  After the
inbound channel is closed and this stage has sent all the values
downstream, it closes the outbound channel:

{{code "pipelines/square.go" `/func sq/` `/^}/`}}

The `main` function sets up the pipeline and runs the final stage: it receives
values from the second stage and prints each one, until the channel is closed:

{{code "pipelines/square.go" `/func main/` `/^}/`}}

Since `sq` has the same type for its inbound and outbound channels, we
can compose it any number of times.  We can also rewrite `main` as a
range loop, like the other stages:

{{code "pipelines/square2.go" `/func main/` `/^}/`}}

## Fan-out, fan-in

Multiple functions can read from the same channel until that channel is closed;
this is called _fan-out_. This provides a way to distribute work amongst a group
of workers to parallelize CPU use and I/O.

A function can read from multiple inputs and proceed until all are closed by
multiplexing the input channels onto a single channel that's closed when all the
inputs are closed.  This is called _fan-in_.

We can change our pipeline to run two instances of `sq`, each reading from the
same input channel.  We introduce a new function, _merge_, to fan in the
results:

{{code "pipelines/sqfan.go" `/func main/` `/^}/`}}

The `merge` function converts a list of channels to a single channel by starting
a goroutine for each inbound channel that copies the values to the sole outbound
channel.  Once all the `output` goroutines have been started, `merge` starts one
more goroutine to close the outbound channel after all sends on that channel are
done.

Sends on a closed channel panic, so it's important to ensure all sends
are done before calling close.  The
[`sync.WaitGroup`](/pkg/sync/#WaitGroup) type
provides a simple way to arrange this synchronization:

{{code "pipelines/sqfan.go" `/func merge/` `/^}/`}}

## Stopping short

There is a pattern to our pipeline functions:

  - stages close their outbound channels when all the send operations are done.
  - stages keep receiving values from inbound channels until those channels are closed.

This pattern allows each receiving stage to be written as a `range` loop and
ensures that all goroutines exit once all values have been successfully sent
downstream.

But in real pipelines, stages don't always receive all the inbound
values.  Sometimes this is by design: the receiver may only need a
subset of values to make progress.  More often, a stage exits early
because an inbound value represents an error in an earlier stage. In
either case the receiver should not have to wait for the remaining
values to arrive, and we want earlier stages to stop producing values
that later stages don't need.

In our example pipeline, if a stage fails to consume all the inbound values, the
goroutines attempting to send those values will block indefinitely:

{{code "pipelines/sqleak.go" `/first value/` `/^}/`}}

This is a resource leak: goroutines consume memory and runtime resources, and
heap references in goroutine stacks keep data from being garbage collected.
Goroutines are not garbage collected; they must exit on their own.

We need to arrange for the upstream stages of our pipeline to exit even when the
downstream stages fail to receive all the inbound values.  One way to do this is
to change the outbound channels to have a buffer.  A buffer can hold a fixed
number of values; send operations complete immediately if there's room in the
buffer:

{{raw `
	c := make(chan int, 2) // buffer size 2
	c <- 1  // succeeds immediately
	c <- 2  // succeeds immediately
	c <- 3  // blocks until another goroutine does <-c and receives 1
`}}

When the number of values to be sent is known at channel creation time, a buffer
can simplify the code.  For example, we can rewrite `gen` to copy the list of
integers into a buffered channel and avoid creating a new goroutine:

{{code "pipelines/sqbuffer.go" `/func gen/` `/^}/`}}

Returning to the blocked goroutines in our pipeline, we might consider adding a
buffer to the outbound channel returned by `merge`:

{{code "pipelines/sqbuffer.go" `/func merge/` `/unchanged/`}}

While this fixes the blocked goroutine in this program, this is bad code.  The
choice of buffer size of 1 here depends on knowing the number of values `merge`
will receive and the number of values downstream stages will consume.  This is
fragile: if we pass an additional value to `gen`, or if the downstream stage
reads any fewer values, we will again have blocked goroutines.

Instead, we need to provide a way for downstream stages to indicate to the
senders that they will stop accepting input.

## Explicit cancellation

When `main` decides to exit without receiving all the values from
`out`, it must tell the goroutines in the upstream stages to abandon
the values they're trying to send.  It does so by sending values on a
channel called `done`.  It sends two values since there are
potentially two blocked senders:

{{code "pipelines/sqdone1.go" `/func main/` `/^}/`}}

The sending goroutines replace their send operation with a `select` statement
that proceeds either when the send on `out` happens or when they receive a value
from `done`.  The value type of `done` is the empty struct because the value
doesn't matter: it is the receive event that indicates the send on `out` should
be abandoned.  The `output` goroutines continue looping on their inbound
channel, `c`, so the upstream stages are not blocked. (We'll discuss in a moment
how to allow this loop to return early.)

{{code "pipelines/sqdone1.go" `/func merge/` `/unchanged/`}}

This approach has a problem: _each_ downstream receiver needs to know the number
of potentially blocked upstream senders and arrange to signal those senders on
early return.  Keeping track of these counts is tedious and error-prone.

We need a way to tell an unknown and unbounded number of goroutines to
stop sending their values downstream.  In Go, we can do this by
closing a channel, because
[a receive operation on a closed channel can always proceed immediately, yielding the element type's zero value.](/ref/spec#Receive_operator)

This means that `main` can unblock all the senders simply by closing
the `done` channel.  This close is effectively a broadcast signal to
the senders.  We extend _each_ of our pipeline functions to accept
`done` as a parameter and arrange for the close to happen via a
`defer` statement, so that all return paths from `main` will signal
the pipeline stages to exit.

{{code "pipelines/sqdone3.go" `/func main/` `/^}/`}}

Each of our pipeline stages is now free to return as soon as `done` is closed.
The `output` routine in `merge` can return without draining its inbound channel,
since it knows the upstream sender, `sq`, will stop attempting to send when
`done` is closed.  `output` ensures `wg.Done` is called on all return paths via
a `defer` statement:

{{code "pipelines/sqdone3.go" `/func merge/` `/unchanged/`}}

Similarly, `sq` can return as soon as `done` is closed.  `sq` ensures its `out`
channel is closed on all return paths via a `defer` statement:

{{code "pipelines/sqdone3.go" `/func sq/` `/^}/`}}

Here are the guidelines for pipeline construction:

  - stages close their outbound channels when all the send operations are done.
  - stages keep receiving values from inbound channels until those channels are closed or the senders are unblocked.

Pipelines unblock senders either by ensuring there's enough buffer for all the
values that are sent or by explicitly signalling senders when the receiver may
abandon the channel.

## Digesting a tree

Let's consider a more realistic pipeline.

MD5 is a message-digest algorithm that's useful as a file checksum.  The command
line utility `md5sum` prints digest values for a list of files.

	% md5sum *.go
	d47c2bbc28298ca9befdfbc5d3aa4e65  bounded.go
	ee869afd31f83cbb2d10ee81b2b831dc  parallel.go
	b88175e65fdcbc01ac08aaf1fd9b5e96  serial.go

Our example program is like `md5sum` but instead takes a single directory as an
argument and prints the digest values for each regular file under that
directory, sorted by path name.

	% go run serial.go .
	d47c2bbc28298ca9befdfbc5d3aa4e65  bounded.go
	ee869afd31f83cbb2d10ee81b2b831dc  parallel.go
	b88175e65fdcbc01ac08aaf1fd9b5e96  serial.go

The main function of our program invokes a helper function `MD5All`, which
returns a map from path name to digest value, then sorts and prints the results:

{{code "pipelines/serial.go" `/func main/` `/^}/`}}

The `MD5All` function is the focus of our discussion.  In
[serial.go](pipelines/serial.go), the implementation uses no concurrency and
simply reads and sums each file as it walks the tree.

{{code "pipelines/serial.go" `/MD5All/` `/^}/`}}

## Parallel digestion

In [parallel.go](pipelines/parallel.go), we split `MD5All` into a two-stage
pipeline.  The first stage, `sumFiles`, walks the tree, digests each file in
a new goroutine, and sends the results on a channel with value type `result`:

{{code "pipelines/parallel.go" `/type result/` `/}/` "HLresult"}}

`sumFiles` returns two channels: one for the `results` and another for the error
returned by `filepath.Walk`.  The walk function starts a new goroutine to
process each regular file, then checks `done`.  If `done` is closed, the walk
stops immediately:

{{code "pipelines/parallel.go" `/func sumFiles/` `/^}/`}}

`MD5All` receives the digest values from `c`.  `MD5All` returns early on error,
closing `done` via a `defer`:

{{code "pipelines/parallel.go" `/func MD5All/` `/^}/` "HLdone"}}

## Bounded parallelism

The `MD5All` implementation in [parallel.go](pipelines/parallel.go)
starts a new goroutine for each file. In a directory with many large
files, this may allocate more memory than is available on the machine.

We can limit these allocations by bounding the number of files read in
parallel.  In [bounded.go](pipelines/bounded.go), we do this by
creating a fixed number of goroutines for reading files.  Our pipeline
now has three stages: walk the tree, read and digest the files, and
collect the digests.

The first stage, `walkFiles`, emits the paths of regular files in the tree:

{{code "pipelines/bounded.go" `/func walkFiles/` `/^}/`}}

The middle stage starts a fixed number of `digester` goroutines that receive
file names from `paths` and send `results` on channel `c`:

{{code "pipelines/bounded.go" `/func digester/` `/^}/` "HLpaths"}}

Unlike our previous examples, `digester` does not close its output channel, as
multiple goroutines are sending on a shared channel.  Instead, code in `MD5All`
arranges for the channel to be closed when all the `digesters` are done:

{{code "pipelines/bounded.go" `/fixed number/` `/End of pipeline/` "HLc"}}

We could instead have each digester create and return its own output
channel, but then we would need additional goroutines to fan-in the
results.

The final stage receives all the `results` from `c` then checks the
error from `errc`.  This check cannot happen any earlier, since before
this point, `walkFiles` may block sending values downstream:

{{code "pipelines/bounded.go" `/m := make/` `/^}/` "HLerrc"}}

## Conclusion

This article has presented techniques for constructing streaming data pipelines
in Go.  Dealing with failures in such pipelines is tricky, since each stage in
the pipeline may block attempting to send values downstream, and the downstream
stages may no longer care about the incoming data.  We showed how closing a
channel can broadcast a "done" signal to all the goroutines started by a
pipeline and defined guidelines for constructing pipelines correctly.

Further reading:

  - [Go Concurrency Patterns](/talks/2012/concurrency.slide#1)
    ([video](https://www.youtube.com/watch?v=f6kdp27TYZs)) presents the basics
    of Go's concurrency primitives and several ways to apply them.
  - [Advanced Go Concurrency Patterns](/blog/advanced-go-concurrency-patterns)
    ([video](http://www.youtube.com/watch?v=QDDwwePbDtw)) covers more complex
    uses of Go's primitives,
    especially `select`.
  - Douglas McIlroy's paper [Squinting at Power Series](https://swtch.com/~rsc/thread/squint.pdf)
    shows how Go-like concurrency provides elegant support for complex calculations.
