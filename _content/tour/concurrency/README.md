# `go function(argumentsToPass)` — goroutine —
* := thread lightweight / managed by Go runtime
  * ⚠️ run in the same address space ⚠️ -- TODO: How to check?
    * → access to shared memory — must be — synchronized
* `argumentsToPass` are evaluated in the 👁️ current 👁️ goroutine
* `function(argumentsToPass)` is executed in the 👁️ NEW goroutine 👁️
* `go run goroutines.go`
