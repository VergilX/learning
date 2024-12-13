package main

import (
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()  // one to many (multiplexer by regex)

    mux.HandleFunc("GET /{$}", home)  // Adds handle "/"-> func home()
    mux.HandleFunc("GET /snippet/view/{id}", snippetView)
    mux.HandleFunc("GET /snippet/create", snippetCreate)
    mux.HandleFunc("POST /snippet/create", snippetCreatePost)

    log.Print("starting server on :4000")

    err := http.ListenAndServe(":4000", mux)  // listens at port ("host:port")
    // if host not mentioned, goes an all network interfaces avaiable

    log.Fatal(err)
}
