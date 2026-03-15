package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux
	// then register the home function as the handler for the
	// "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Print a log message to say that the server is starting
	log.Print("starting server on :4000")

	// Use the http.ListenAndServe() function to start a new web server.
	// We pass in two parameters
	// 1. the TCP network address to listen on (:4000)
	// 2. the servemux we just created.
	// if http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and terminate
	// the program.
	// Note than any error return by http.ListenAndServe() is always non-nil
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
