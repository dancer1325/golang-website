//go:build ignore

package main

/*
specific arguments because it's necessary for http.HandleFunc()
r *http.Request		client HTTP request
*/
func handler(w http.ResponseWriter, r *http.Request) {
	// write to w
	// %s 	pass a string
	// r.URL.Path[1:]		sub-slice of Path  [1@ character, endWithoutTakingInAccount)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	// register a handler for '/'
	http.HandleFunc("/", handler)

	//http.ListenAndServe()
	// 1. listens on a port
	// 2. until the program is terminated -> function is blocked
	// log.Fatal  wrapping to log that error
	log.Fatal(http.ListenAndServe(":8080", nil))
}
