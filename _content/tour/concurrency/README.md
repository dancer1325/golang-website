# `go function(argumentsToPass)` — goroutine —
* := thread lightweight / managed by Go runtime
  * ⚠️ run in the same address space ⚠️ -- TODO: How to check?
    * → access to shared memory — must be — synchronized
* `argumentsToPass` are evaluated in the 👁️ current 👁️ goroutine
* `function(argumentsToPass)` is executed in the 👁️ NEW goroutine 👁️
* `go run goroutines.go`

# `chan` — channel —
* := typed conduit which
  * `channelVariable <- valueToSend`
    * the channel receive valueS
  * `assignToAVariable, boolVariable <- channelVariable`
    * the channel send valueS 
    * 👁️the goroutine with this statement waits for receiving it 👁️
    * `boolVariable`
      * if `true` == channel can send more values
      * if `false` == channel is closed == NO more values will be sent
* ways to create them
  * `make(chan TypesToSenViaTheChannel)`
* uses
  * sync goroutines without explicit locks OR condition variables -- run several times the next command, to check random different order printing the new goroutines--
* `close(channelVariable)`
  * == close the channel == NO more values will be sent
  * 👁️ JUST the sender should close the channel 👁️
  * ⚠️ if you send values to a closed channel → panic will be thrown ⚠️
  * uses
    * 👁️ receiver must be told that NO more expected values 👁️ -- Example: `range` --
* `for variableDeclaration := range channelVariable {…}`
  * infinite loop till channel is closed
* `go run channels.go` & `go run range-and-close.go`

# buffered channels
* := channel with a defined length
  * 👁️previous to send values the channel ALL channel must be fulfilled 👁️
  * 👁️ONCE you read the values from the channel → channel is empty👁️
* ways to create them
  * `make(chan TypesToSenViaTheChannel, bufferLength)`
* `go run buffered-channels.go`
