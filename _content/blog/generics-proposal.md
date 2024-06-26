---
title: A Proposal for Adding Generics to Go
date: 2021-01-12
by:
- Ian Lance Taylor
tags:
- go2
- proposals
- generics
summary: Generics is entering the language change proposal process
---

## Generics proposal

We’ve filed [a Go language change
proposal](/issue/43651) to add support for type
parameters for types and functions, permitting a form of generic
programming.

## Why generics?

Generics can give us powerful building blocks that let us share code
and build programs more easily.
Generic programming means writing functions and data structures where
some types are left to be specified later.
For example, you can write a function that operates on a slice of some
arbitrary data type, where the actual data type is only specified when
the function is called.
Or, you can define a data structure that stores values of any type,
where the actual type to be stored is specified when you create an
instance of the data structure.

Since Go was first released in 2009, support for generics has been one
of the most commonly requested language features.
You can read more about why generics are useful in 
[an earlier blog post](/blog/why-generics).

Although generics have clear use cases, fitting them cleanly into a
language like Go is a difficult task.
One of the [first (flawed) attempts to add generics to
Go](/design/15292/2010-06-type-functions) dates back
all the way to 2010.
There have been several others over the last decade.

For the last couple of years we’ve been working on a series of design
drafts that have culminated in [a design based on type
parameters](/design/go2draft-type-parameters).
This design draft has had a lot of input from the Go programming
community, and many people have experimented with it using the
[generics playground](https://go2goplay.golang.org) described in [an
earlier blog post](/blog/generics-next-step).
Ian Lance Taylor gave [a talk at GopherCon
2019](https://www.youtube.com/watch?v=WzgLqE-3IhY)
about why to add generics and the strategy we are now following.
Robert Griesemer gave [a follow-up talk about changes in the design,
and the implementation, at GopherCon
2020](https://www.youtube.com/watch?v=TborQFPY2IM).
The language changes are fully backward compatible, so existing Go
programs will continue to work exactly as they do today.
We have reached the point where we think that the design draft is good
enough, and simple enough, to propose adding it to Go.

## What happens now?

The [language change proposal process](/s/proposal)
is how we make changes to the Go language.
We have now [started this process](/issue/43651)
to add generics to a future version of Go.
We invite substantive criticisms and comments, but please try to avoid
repeating earlier comments, and please try to [avoid simple plus-one
and minus-one comments](/wiki/NoPlusOne).
Instead, add thumbs-up/thumbs-down emoji reactions to comments with
which you agree or disagree, or to the proposal as a whole.

As with all language change proposals, our goal is to drive toward a
consensus to either add generics to the language or let the proposal
drop.
We understand that for a change of this magnitude it will be
impossible to make everybody in the Go community happy, but we intend
to get to a decision that everybody is willing to accept.

If the proposal is accepted, our goal will be to have a complete,
though perhaps not fully optimized, implementation for people to try
by the end of the year, perhaps as part of the Go 1.18 betas.
