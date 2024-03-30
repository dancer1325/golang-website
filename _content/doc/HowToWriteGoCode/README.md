* https://go.dev/doc/code

# TODO:

# Testing
* lightweight test framework
  * == `go test` + [testing](https://pkg.go.dev/testing@go1.22.1)
  * run each function
  * if the function calls a failure function → test is considered as failed
    * *Example of failure function:* `t.Error` OR `t.Fail`
* how to create?
  * file
    * name —  `*_test.go` —
    * with functions
      * name — `TestXXX` —
      * signature — `(t *testing.T)` —
* how to run?
  * `go test`

# Notes
* NOT allowed creating different packages in the same directory
  * -> 'morestrings' was created in a subdirectory