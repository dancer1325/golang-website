# Methods
* NOT classes
* methods on types can be defined
* := function with special receiver argument
  * `func receiver methodName.. {}`
* `go run methods.go`

# TODO:

# `type interfaceName interface` -- interface type --
* := set of method signatures
* ⚠️ if a variable implements interface’s methods → that interface type can hold it ⚠️
  * 🧠 although at run time the value stored in the interface variable could change type, statically it’s typed 🧠
* `go run intefaces.go`

# `interface{}` -- empty interface - 
* := `interface` / NO methods
  * 👁️can hold values of any type👁️
    * **Reason:** 🧠 every type implements zero methods🧠
* uses
  * code which handles values of unknown type
* `go run empty-interface.go`

# TODO: