package main

import (
	"fmt"
	"net/http"
)

func main() {
	//the HandleFunc method accepts a path and a function as arguments

	http.HandleFunc("/", handler)

	//after defining our server, we finally "listen and serve" on port 8080
	//the second argument is the handler, which we now left as nil
	//and the handler defined above (in HandleFunc) is used
	http.ListenAndServe(":8080", nil)
}

// handler is our handler function. It has to follow the function signature
// of a ResponseWriter and Request type as the arguments.
func handler(w http.ResponseWriter, r *http.Request) {

	//for this case we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World")
	//Fprintf takes a "writer" As its first argument.
	//the second argument is the data that is piped into this writer.
	//the output therefore appears according to where the writermoves it.
	//in this cse the ResponseWriter w writes the output as the response to the users request.

}
