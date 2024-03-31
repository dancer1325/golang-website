* https://www.youtube.com/watch?v=h6Cw9iCDVcU

* Purpose
  * ⭐create alternate equivalent names — for — types ⭐
    * → gradual code repair  — during — codebase refactoring
* Go’s goal
  * build software which scales
    * scale the size of systems / built with Go
      * == # of computers OR amount of data to process…
    * scale the size of Go programs
      * == large codebases + large # of engineers + large # of changes …

# Codebase refactoring
* := process of rethinking & reviewing decisions about
  * code grouped under package
  * package1 ← relationship with → package2
* Reasons
  * package — is split into → > 1 package
    * Example: regular expression parser used by specific users → another exported package https://pkg.go.dev/regexp/syntax — ≠ https://pkg.go.dev/regexp — 
  * improve naming
  * lighten dependencies
  * change dependency graph
    * Example: `os.Error` was used by many packages → `os` package was imported by many others. Solution: new package `errors` was created

# Gradual code repair
* 3 stages
  * pattern for stage 2
    * constants
      * TODO:
    * functions
      * TODO:
    * types
      * NO way
        * →
          * use Conventional code repair OR
          * next suggested solutions
        * Example1: Go’s `os.Error` → `error`
          * 👁typeName1 ← can NOT be exchangeable with → typeName2 👁️
        * Example2: Kubernetes split out the `IntOrString type` under the package util → package intstr
          * Reason to apply gradual repair: it’s used as `struct`'s field
        * TODO:
  * use cases
    * 👁 infeasible to do all the repairs at once 👁️
    * code to be repaired — is spread across — multiple repositories

# Conventional code repair
* == 1! commit for all
* pros
  * NOT break in the codebase
* cons
  * if large codebases → 
    * 1 big commit
    * while you are preparing it → another guy added functionality

# Suggested solutions
* alias
  * https://go.googlesource.com/proposal/+/master/design/16339-alias-decls.md
  * ==  uniform way to create an alternate name / 👁️any kind of identifier 👁️
* versioning
  * TODO:
* type aliases
  * ==  uniform way to create an alternate name / 👁️just type 👁️
    * ≠ alias
  * `type OldAPI = NewPackage.API`
