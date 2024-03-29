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

# interface values
* == `(value, type)`
* if you call an interface value’s method → underlying `type`'s method (with same name) is executed
* `go run interface-values.go`

# `interface{}` -- empty interface - 
* := `interface` / NO methods
  * 👁️can hold values of any type👁️
    * **Reason:** 🧠 every type implements zero methods🧠
* uses
  * code which handles values of unknown type
* `go run empty-interface.go`

# `:= interfaceName.(underlyingInterfaceValueType)` -- Type assertions --
* `:= interfaceName.(underlyingInterfaceValueType)`
  * syntax
  * returns `interfaceValue’svalue, boolean`
    * syntax == mutating maps
    * boolean is optional
    * if you pass wrong interface value’s type &
      * boolean passed to return → boolean is false & NOT panic occurs
      * NOT boolean passed to return → panic occurs
* allows
  * getting access to interface value's value
* `go run type-assertions.go`

# TODO: