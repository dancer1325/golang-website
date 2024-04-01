# `for initStatement; condition; postStatement {...}`
* syntax
  * NO `()`
  * `{}` always required
* ⚠️ 1! looping construct in Go ⚠️
  * ->
    * is Go’s `while`
* `initStatement`
  * before first iteration → it’s executed
  * normally, it’s a short variable declaration
    * these variables are visible in the `for`'s scope
  * 👁️ Optional 👁️
* `condition`
  * before each iteration → it’s executed
  * 👁️ Mandatory 👁️
    * → if you omit it → forever loop
* `postStatement`
  * at the end of each iteration → it’s executed
  * 👁️ Optional 👁️
* `go run for.go` & `go run for-continued.go` & `go run for-is-gos-while.go` & `go run forever.go`