* Principles of Golang
* Inspired by the proverbs done for the [Go game](https://en.wikipedia.org/wiki/Go_(game))

# Don't communicate by sharing memory, share memory by communicating
* '../codewalk/Share Memory by Communicating'

# Concurrency is not parallelism
* Check '../2012/Concurrency!=Parallelism'

# Channels orchestrate; mutexes serialize.
* mutex
  * fine-grained & very small
  * tend to -> serialize execution
    * if you mutex a variable -> 1! thing can happen in that variable / time

# The bigger the interface, the weaker the abstraction
* `io.Reader` & `io.Writer` are the most important types and just have 2/3 methods

# Make the zero value useful
* Check '../tour/basics'
* allows
  * avoid using constructor

# interface `{}` says nothing
* Check '../tour/methods'

# Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.
* Check '../blog/gofmt'

# A little copying is better than a little dependency.
* Keep dependency tree tiny -> programs
  * compile faster
  * easier to maintain
  * simpler

# Syscall must always be guarded with build tags.
* Check 'std/syscall'

# Cgo must always be guarded with build tags.
* Check '../blog/cgo'
* Reason == Reason for syscall
  * NOT portable

# Cgo is not Go.

# With the unsafe package there are no guarantees.
* Check '/std/unsafe'
 
# Clear is better than clever.
* For learner -> better to write simple code
 
# Reflection is never clear.
* Check '/blog/The Laws of Reflection'
* In runtime checks

# Errors are values.
* Check '../blog/Errors are values'

# Don't just check errors, handle them gracefully.
* Such as
  * enrich
  * make something before logging
 
# Design the architecture, name the components, document the details.

# Documentation is for users.
* From the user perspective

# Don't panic.
