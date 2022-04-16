package main

import (
	"net/http"
)

//this function writes a byte slice containing "Hello from SnippetBox" as the response body
//this is the home handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetBox"))
}

func main() {
	//initialize a new servemux and register "home" as the handler for the "/" URL
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

}
