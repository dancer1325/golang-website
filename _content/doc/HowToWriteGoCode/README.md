* https://go.dev/doc/code

# Goal
* develop a Go package inside a module
* Check golang repo 'standardLibrary/cmd/go'

# Package
* == way to organize Go programs
* := collection of source files in the same directory which  -- Check in this directory how all contains `package main` --
  * 👁️ compiled together 👁️
  * functions & types & variables & constants in a source file 👁️— are visible to — rest of package’s files 👁️ -- TODO: Fix how --
* `import “ModulePathAndSubdirectoryWithinModule”` 
  * allows importing a package
* `go run main2.go`

# Module — “go.mod” —
* ≥ 1 module / repository
  * **Note:** 👁️Normally there is 1! module / repository located at root of the repository 👁️ -- Check in this project -- 
* := collection of related Go packages / released together
  * also contains package’s packages
* `module path`
  * `path`
    * 👁️ import path prefix 👁️ for all module’s packages -- Check 'https://pkg.go.dev/golang.org/x/website' for this project --
      * 👁️ @Standard Library ‘s packages do NOT contain the prefix 👁️
    * → indicates where to download it

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
