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
            * Reason: There is NOT yet built handler for '/save/'
* `go run final-template.go` OR `go build final-template.go` & `./final-template` + allowed pop up panel
  * save a .txt file the information passed via the form in '/edit/'
  * Open in your desired browser or trigger the curl
    * 'localhost:8080/edit/createdViaEditForm'    -- file created via edit form --
* `go run final-noclosure.go` OR `go build final-noclosure.go` & `./final-noclosure` + allowed pop up panel
  * `http.Error()` to handle errors
  * Open in your desired browser or trigger the curl
    * Problems: Which request OR how to force?
      * Attempt1: Comment couple of lines in viewHandler
      * Attempt2: Check [dlv](https://github.com/go-delve/delve/tree/master/Documentation/installation)
      * Solution: TODO:

# html/template
* allows
  * HTML file (_Example:_ 'edit.html') -- is separated from -- Go code
    * '>'  -- is replaced by '&gt'
      * Problems: How to add commentaries to the template 
        * Solution: `{{/* Commentary */}}`
* `go run final-noerror.go` OR `go build final-noerror.go` & `./final-noerror` + allowed pop up panel
  * replace hardcoded HTML by path to them
  * steps
    * same as previous one
* `go run part3-errorhandling.go` OR `go build part3-errorhandling.go` & `./part3-errorhandling` + allowed pop up panel
  * if the page '/view/*' to visualize does NOT exist -> redirect to '/edit/*'
  * Open in your desired browser or trigger the curl
    * 'localhost:8080/view/NoExistingPage' -- is redirected to -> 'localhost:8080/edit/NoExistingPage' 

