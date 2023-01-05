package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	//here we form a new HTTP request
	//this is the request that's going to be passed to our handler
	//the first argument is the method, the second argument is the route,
	//which we blank for now, and the third argument is the request body,
	//which we don't have in this case
	req, err := http.NewRequest("GET", "", nil)

	//in case there is an error in forming the request, we fail and stop the test
	if err != nil {
		t.Fatal(err)
	}

	//we use GO's httptest library to create an http recorder.
	//this recorder will act as the target of our http request
	//you can think of it as a mini-browser, which will accept
	//the result of the http request that we make.
	recorder := httptest.NewRecorder()

	//create an HTTP handler from our handler function.
	hf := http.HandlerFunc(handler)

	//serve the HTTP request to our recorder. This is the line
	//that actually executes our handler that we want to test
	hf.ServeHTTP(recorder, req)

	//check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong code: got %v want %v", status, http.StatusOK)
	}

	//check the response body is what we expect.
	expected := `Hello World!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
