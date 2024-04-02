# Pointers
* := types which
  * hold memory address of a value
  * ⚠️ are NOT arithmetic ⚠️
    * Note: ≠ C
  * `*`
    * `*T`
      * type of the pointer, whose type of the value is `T`
    * `*pointerValue` = `value` -- "dereferencing" OR "indirecting" —
  * `&i`
    * operand `i` — generates a → pointer of `i`
  * `nil`
    * zero value of a pointer
* `go run pointers.go`

# `type structName struct { fields }` — structs —
* := collection of fields /
  * ways to access to the fields
    * `structVariable.structField`
    * `(*pointerToStructVariable).structField`
    * `pointerToStructVariable.structField`
      * Note: 👁️ unnecessary to use dereference 👁️
* `&structVariable` — generates a → pointer of `structVariable`
* `go run structs.go` & `go run struct-fields.go` & `go run struct-pointers.go`