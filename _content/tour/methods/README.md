# Methods
* NOT classes
* == methods on types
  * `func receiver methodName.. {}`
    * receiver can be
      * type
        * ⚠️ receiver into the method is a copy → outside the method is NOT modified ⚠️
        * ⚠️ Either value or pointer is valid ⚠️
      * pointer
        * Reasons to use it:
          * receiver must be changed
          * avoid copying the value / each method call
        * ⚠️ Either value or pointer is valid ⚠️
* == function with special receiver argument
* `go run methods.go` & `go run methods-funcs.go` & `go run methods-continued.go` & `go run methods-pointers.go` & `go run indirection.go` & `go run indirection-values.go` 

# Functions
* ‘s arguments can be
  * type
  * pointer
    * ⚠️ argument into the function is a copy → outside the function is NOT modified  ⚠️
    * ⚠️ JUST pointers are accepted ⚠️
* `go run methods-pointers-explained.go` & `go run indirection.go` & `go run indirection-values.go`

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