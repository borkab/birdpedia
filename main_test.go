package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ErrCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	//Most of the code is similar. The only difference is that now we make a
	//request to a route we know we didn't define, like the `POST /hello` route.
	resp, err := http.Post(mockServer.URL+"/hello", "", nil)

	ErrCheck(err)

	//We want our status to be 405(method not allowed)
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	//the code to test the body is also mostly the same, except this time,
	//we expect an empty body
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)

	ErrCheck(err)
	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}

func TestRouter(t *testing.T) {
	//instantiate the router using the constructor function that
	//we defined previously
	r := newRouter()

	//create a new server using the "httptest" libraries
	//`NewServer` method
	//Documentation: https://golang.org/pkg/net/http/httptest/#NewServer
	mockServer := httptest.NewServer(r)

	//the mock server we created runs a server and exposes its
	// location in the URL attribute
	//we make a GET request to thee "hello" route we defined
	//in the router
	resp, err := http.Get(mockServer.URL + "/hello")

	//handle any unexpected error
	ErrCheck(err)

	//we want our status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	//in the next few lines, the response body is read,
	//and converted to a string
	defer resp.Body.Close()
	//read the body into a bunch of bytes (b)
	b, err := io.ReadAll(resp.Body)

	ErrCheck(err)

	//convert the bytes to a string
	respString := string(b)
	expected := "Hello World!"

	//we want our response to match the one defined in our handler
	//if it does happen to be "Hello World!", then it confirms,
	// that the route is correct
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestHandler(t *testing.T) {
	//here we form a new HTTP request
	//this is the request that's going to be passed to our handler
	//the first argument is the method, the second argument is the route,
	//which we blank for now, and the third argument is the request body,
	//which we don't have in this case
	req, err := http.NewRequest("GET", "", nil)

	//in case there is an error in forming the request, we fail and stop the test
	ErrCheck(err)

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
