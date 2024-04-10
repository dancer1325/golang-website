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

# `type interfaceName interface` -- interface type --
* := set of method signatures
* ⚠️ if a variable implements interface’s methods → that interface type can hold it ⚠️
  * 🧠 although at run time the value stored in the interface variable could change type, statically it’s typed 🧠
* ⚠️ if a type implements interface’s methods → that interface is implemented ⚠️
  * NOT `implements` keyword exist
  * 👁️ implicit implementation 👁️
    * == implementation can be delegated to another package
* `go run intefaces.go` & `go run interfaces-are-satisfied-implictly.go`

# interface values
* == `(value, type)`
* if you call an interface value’s method → underlying `type`'s method (with same name) is executed
* if the `value` is nil → method called with a nil receiver ⚠️ WITHOUT triggering null pointer ⚠️
* nil interface
  * NOT hold `value` NOR `type`
  * if you call a method → Null Pointer error in runtime
* `go run interface-values.go` & `go run interface-values-with-nil.go` & `go run nil-interface-values.go`

# `interface{}` -- empty interface - 
* := `interface` / NO methods
  * 👁️can hold values of any type👁️
    * **Reason:** 🧠 every type implements zero methods🧠
* uses
  * code which handles values of unknown type
* `go run empty-interface.go`

# `:= interfaceName.(underlyingInterfaceValueType)` -- Type assertions --
* returns `interfaceValue’svalue, boolean`
  * syntax == mutating maps
  * boolean is optional
  * if you pass wrong interface value’s type &
    * boolean passed to return → boolean is false & NOT panic occurs
    * NOT boolean passed to return → panic occurs
* allows
  * getting access to interface value's value
* `go run type-assertions.go`

# `switch initStatement := interfaceName.(type) { case type1: … case type2: ... default: ...}` — type switches —
* `interfaceName.(type)` == type assertion BUT `type` is a keyword
* == switch statement BUT specifying types (NOT values)
* if there is NO match with any case == `default` → initStatement’s type is `type`
* allows
  * several type assertions in series
* `go run type-switches.go`

# TODO: