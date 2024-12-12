package main

import (
    "log"
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from SnippetBox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
    mux := http.NewServeMux()  // one to many (multiplexer by regex)

    mux.HandleFunc("/{$}", home)  // Adds handle "/"-> func home()
    mux.HandleFunc("/snippet/view", snippetView)  
    mux.HandleFunc("/snippet/create", snippetCreate)  

    log.Print("starting server on :4000")

    err := http.ListenAndServe(":4000", mux)  // listens at port ("host:port")
    // if host not mentioned, goes an all network interfaces avaiable

    log.Fatal(err)
}
