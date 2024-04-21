* https://go.dev/doc/articles/wiki/

# Goal
* create data structure with load and save methods
* use
  * 'net/http' -- for build -- web applications
  * 'html/template' -- to process -- HTML templates
  * 'regexp' -- to validate -- user input
  * closures

# Create data structure with load and save methods
* `go run part1-noerror.go` OR `go build part1-noerror.go` & `./part1-noerror`
  * 'loadPage' does NOT handle the error
  * creates a file named 'TestPage.txt' in this directory
  * Problems
    * Problem1: "undefined: fmt"
      * Solution: `import "fmt"`
* `go run part1.go` OR `go build part1.go` & `./part1`
  * 'loadPage' handles the error
  * creates a file named 'TestPage.txt' in this directory

# net/http
* `go run http-sample.go` OR `go build http-sample.go` & `./http-sample` + allowed pop up panel
  * Open in your desired browser or trigger the curls
    * 'localhost:8080'    -- response without url path --
    * 'localhost:8080/Alfred'    -- response adding the url path 'Alfred' --
* `go run part2.go` OR `go build part2.go` & `./part2` + allowed pop up panel
  * steps
    * create a file manually "NameToLoadAndDisplay.txt"
    * Open in your desired browser or trigger the curl
      * 'localhost:8080/view/NameToLoadAndDisplay'    -- content of the file is displayed!! --
* `go run notemplate.go` OR `go build notemplate.go` & `./notemplate` + allowed pop up panel
  * hardcoded HTML without templating 
  * steps
    * same as previous one
    * Open in your desired browser or trigger the curl
      * 'localhost:8080/edit/NameToLoadAndDisplay'    -- content of the file is displayed!! --
        * Problems:
          * Problem1: Why content of the file NOT updated after clicking in save button ?
            * Attempt1: `save`
            * Attempt2: `save()`
            * Solution: TODO:

