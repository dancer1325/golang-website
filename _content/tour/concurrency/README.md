# `go function(argumentsToPass)` — goroutine —
* := thread lightweight / managed by Go runtime
  * ⚠️ run in the same address space ⚠️ -- TODO: How to check?
    * → access to shared memory — must be — synchronized
* `argumentsToPass` are evaluated in the 👁️ current 👁️ goroutine
* `function(argumentsToPass)` is executed in the 👁️ NEW goroutine 👁️
* `go run goroutines.go`

# `chan` — channel —
* := typed conduit which
  * receive valueS the channel -- via `channelVariable <- valueToSend`
  * send valueS the channel -- via `assignToAVariable <- channelVariable`
    * 👁️the goroutine with this statement waits for receiving it 👁️
* ways to create them
  * `make(chan TypesToSenViaTheChannel)`
* uses
  * sync goroutines without explicit locks OR condition variables -- run several times the next command, to check random different order printing the new goroutines-- 
* `go run channels.go`

# buffered channels
* Check previous section
  * 👁️previous to send values the channel ALL channel must be fulfilled 👁️
  * 👁️ONCE you read the values from the channel → channel is empty👁️
* ways to create them
  * `make(chan TypesToSenViaTheChannel, bufferLength)`
* `go run buffered-channels.go`
