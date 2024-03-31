# Packages
* ⭐ Go program
  *  == Packages ⭐
  * starts running the main package
* (By convention) name == last element of `import`   -- Check packages.go --
* different can be added in the same directory  -- Check all are main + impossible --
* `go run packages.go`

# Imports
* `import (…)`
  * “factor” all the imports
  * path reference to packages can NOT contain whitespaces
* — can only refer to — exported names -- Check exported names section --
* `go run imports.go`

# Exported names
* requirement
  * 👁️ begins with a capital letter 👁️
* `go run exported-names.go`

# Functions
* — accepts — ≥ 0 arguments
* `func functionName(argumentVarible argumentType, …) returnedType {…}`
  * `argumentVariable argumentType` ≠ C `argumentType argumentVariable`
  * if several consecutive arguments share the type → argumentType can be specified for the last one -- 
  `func functionName(argumentVariable1, argumentVariable2 argumentType, …) returnedType {…}`--
  * if you pass more argumentVariables than defined in the function → errors in Runtime just
  * `(returnedType1, returnedType2, …)` is feasible
  * `(returnedVariable1 returnedType1, returnedVariable2 returnedType2, …)`
    * named return values
      * uses
        * document the meaning
    * == variables declared as top of the function
    * “naked” return — `return` —
      * NOT arguments passed
      * returns the named return values
      * uses
        * short functions 
* `go run functions.go` & `go run functions-continued.go` & `go run multiple-results.go` & `go run named-results.go`

# Variables
* `var variable1, variable2, … variableType`
  * allows
    * declaring list of variables at
      * package level
      * function level
* can be "factored" `()` into -- blocks 
* can be initialized
  * 1 / variable
  * → type can be omitted
* `variable := value` — short declaration —
  * requirements
    * inside a function
* `go run variables.go` & `go run variables-with-initializers.go` & `go run short-variable-declarations.go`

# Basic types
* Are
  * `bool`
  * `string`
  * `int  int8  int16  int32  int64`
  * `uint uint8 uint16 uint32 uint64 uintptr`
  * `byte`
    * alias for `uint8` 
  * `rune`
    * alias for `int32`
    * == Unicode code point (?)
  * `float32 float64`
  * `complex64 complex128`
* Usually 
  * if 32-bit system -> `int uint uintptr` are 32-bit 
  * if 64-bit system -> `int uint uintptr` are 64-bit 
* if you need an integer value -> use `int`
  * sized or unsigned inter types for specific reasons
* `go run variables.go`

# Zero values
* := value given if you declare a variable / no explicit initial value
  * numeric types -> `0`
  * boolean type -> `false`
  * string -> `""`
* `go run zero.go`

# `T(v)` -- Type conversions --
* value `v` — is converted to — type `T`
* 👁️explicit conversion is required 👁️
* `go run type-conversions.go`

TODO: Rest
