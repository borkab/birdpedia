package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//declare a new router
	r := mux.NewRouter()

	//this is where the router is useful, it allows us to declare methods
	//that this path will be valid for
	r.HandleFunc("/hello", handler).Methods("GET")

	//we can then pass our router (after declareing all our routes) to this method
	//where previousrly we were leaving our second argument as nil
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World!")
}
