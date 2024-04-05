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
* `structName { fieldsPassingValues }` — struct literal —
  * create a new struct value / passes the fields values directly
* `go run structs.go` & `go run struct-fields.go` & `go run struct-pointers.go` & `go run struct-literals.go`

# Arrays
* `[n]T`
  * array of `n` values of type `T`
  * 0 is the first element
* 👁️can NOT be resized 👁️
* `[n]T {fieldsPassingValues}` — array literal —
  * create a new array value / passes the items values directly
* `go run array.go`

# Slices
* == arrays + dynamically-sized
* == section of an underlying array
  * Reason: 🧠NOT store any data🧠
  * → if you change slice’s elements → underlying array’s elements are changed → slices / same underlying array will be affected
  * capacity — `cap(sliceVariable)` —
    * := # of elements of the underlying array
    * if you adjust `highestIndex` →
      * slice’s length changes
      * 👁️array NOT change → capacity NOT change 👁️
    * if you adjust `lowestIndex` →
      * slice’s length changes &
      * array changes / drops elements -> capacity changes
* `[]T`
  * slice of values of type `T`
  * 0 is the first element
  * length of the slice — `len(sliceVariable)` —
  * `[][]T` — slices of slices — 
* ways to create slices 
  * `[lowestIndex:highestIndex]`
    * == [lowestIndex, highestIndex)
      * **Note:** 👁️highestIndex is NOT included 👁️
    * by default (== if you do NOT specify) →
      * lowestIndex = 0
      * highestIndex = slice’s length
  * `make([]T, lengthValue, capacityValue)`
    * `capacityValue` optional
      * if you do NOT specify → `capacityValue` = `lengthValue`
* `[]T {fieldsPassingValues}` — slice literals —
  * create a new slice value / passes the items values directly
* zero values is `nil`
* `append()`
* practises
  * more common than array
* `go run slices.go` & `go run slices-pointers.go` & `go run slice-literals.go` & `go run slice-bounds.go` & `go run slice-len-cap.go` & `go run nil-slices.go` & `go run making-slices.go` & `go run slices-of-slice.go` & `go run append.go`

# Range form of for loop
* `for first, second range sliceOrMap {…}`
  * `first`
    * `_` == skip it
  * `second`
    * ways to skip it
      * `_`
      * NOT specifying it directly
  * if you iterate over a slice → 
    * `first` == slice item’s position 
    * `second` == copy of the slice item’s value
* `go run range.go` & `go run range-continued.go`