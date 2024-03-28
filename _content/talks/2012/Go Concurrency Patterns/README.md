# Concurrency
* key to design high performance network services
* goal
  * computer's program structured — as — real world works
    * real world works with independent events happening at the same time
* Go's concurrency primitives
  * are
    * goroutines
    * channels
* allows
  * -- expressing -- concurrent execution
* := composition of independent executing computations
  * → way to structure software
* ≠ parallelism
  * if you have 1! processor → your program can be concurrent, but NOT parallel
* history
  * all ideas come from Communicating Sequential Processes -- Hoare --
## Communication
* Required to have a proper concurrent program
* between goroutines
* done by Channels
## Concurrency patterns
* Generator
  * := function which returns a channel
    * `go run generatorboring.go`
  * → channel — are handled as — services
    * Reason1: 🧠delegation inside the function🧠
    * Reason2: 🧠instances of the service can be created🧠
    * `go run generator2boring.go`  -- you can also check the sync of the channels --
* Multiplexing
  * := function with several channels as inputs, returns another channel
    * → things received from several channels — are passed to → same another channel
    * allows
      * avoid blocking between different channels
    * `go run fainboring.go`

# Goroutines
* := independently executing function
  * launched by `go functionName()` inside the program
* == `> .... &`
* `go run goboring.go`
  * nothing is displayed
    * Reason: main does NOT wait to the goroutine & once main returns -> program exists
* `go run waitgoboring.go`
  * some of the iterations of the goroutine execution are displayed, previously that main returns
  * `main` & `go boring()` are run concurrently
* has its own call stack which
  * as it’s required  -- **Note:** ⚠️≠ thread libraries — you need to define the size — ⚠️ --
    * grow
    * shrink
  * starts very small == cheap
    * **Note:** 👁️ you can find thousands of goroutines running in production👁️
* ≠ thread
  * there could be a program with 1 thread & thousands of goroutines
    * **Reason:** 🧠 Go runtime map multiple goroutines — onto a →  smaller number of operating system threads as needed🧠

# Channel
* := writing to a file descriptor
  * == process in other languages (Example: Erlang)
    * Process := writing to a file by name
  * of a static type
* allows
  * goroutine1 <- is connected with -> goroutine2
    * communication’s Go approach
      * (commonly used) communication by sharing memory — is replaced by → share memory by communicating
      * communication by sharing memory’s approach
        * == blob of memory (logs, variables …) / protected from parallel access
    * connect == communicate + synchronize
      * synchronize — via — blocking operation!!
        * if you read from the channel ( `<- c`) → wait for the value is sent
        * if you send to the channel ( `c <- value`) → wait for the receiver is ready
* Syntax          -- Check 'helpers.go' --
  * `var c chan int`  
    * declare
  * `c = make(chan int)`
    * initialize
  * `c := make(chan int)`
    * declare & initialize at the same time
  * `c <- 1`
    * sends value 1 on the channel
  * `value = <-c`
    * receive the value -- from the -- channel
* `go run changoboring.go`
  * goroutines communicate between them -- via -- channels
  * each time that a communication happens is a block operation
## Buffered channels
* NOT synchronization
  * ==
    * NOT blocking operation
    * once it tries to communicate → keep going for the next line

# Restoring sequencing
* TODO:

# Note
* `go run boring.go`
  * NO concurrent program
* `go run mainboring.go`
  * NO concurrent program

