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

# `if initStatement; condition {…} else {…}`
* syntax
  * NO `()`
  * `{}` always required
* `initStatement`
  * before evaluate the condition → it’s executed
  * normally, it’s a short variable declaration
    * these variables are visible in the `if`'s scope
  * 👁️ Optional 👁️
* `condition`
  * 👁️ Mandatory 👁️
* `else {…}`
  * 👁️ Optional 👁️
  * variables declared in `if` → also available in the `else`
* `go run if.go` & `go run if-with-a-short-statement.go` & `go run if-and-else.go`

# `switch initStatement; condition { case....}`
* == `if else`
* == other languages
  * evaluation from top to bottom
  * ⚠️ except to ⚠️
    * `break` statement is NOT needed
      * **Note:** once logic flows in one of it → NOT run in others
    * switch cases NOT need to be
      * `const`
      * integers
* `initStatement`
  * before evaluate the value → it’s executed
  * normally, it’s a short variable declaration
    * these variables are visible in the `switch`'s scope
  * 👁️ Optional 👁️
* `condition`
  * 👁️Optional 👁️
    * if it’s missing == `switch true` → first case enters
* `go run switch.go` & `go run switch-evaluation-order.go` & `go run switch-with-no-condition.go`
