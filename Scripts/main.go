package main

import (
	"log"
	"net/http"
)

//this function writes a byte slice containing "Hello from SnippetBox" as the response body
//this is the home handler function

func home(w http.ResponseWriter, r *http.Request) {

	//ResponseWriter is for assembling an HTTP response and sending to the user
	//Request is a struct which holds info about the current request.
	if r.URL.Path != "/" {
		http.NotFound(w, r) //if the URL is a path that doesn't exist, issue the 404 not found error message and return.
		return
	}
	w.Write([]byte("Hello from SnippetBox"))

}

func showSnippet(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Displaying a snippet..."))

}

func createSnippet(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Creating a new snippet..."))

}

func main() {

	//initialize a new servemux and register "home" as the handler for the "/" URL

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//Use http.ListenAndServe() to start a new web server
	//Two parameters: TCP network address to listen on, and the servemux
	//If it returns an error, use log.Fatal() to record the error

	log.Println("Starting server on :4000")  //The TCP address
	err := http.ListenAndServe(":4000", mux) //Starting the web server
	log.Fatal(err)                           //Log error if there is any

}
