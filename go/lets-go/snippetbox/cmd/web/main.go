package main

import (
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()  // one to many (multiplexer by regex)

    // Create a fileserver for static files
    fileserver := http.FileServer(http.Dir("./ui/static/"))

    // Handler for fileserver
    // all URLs that start with /static/
    // Arg taken is stripped of /static using http.StripPrefix
    mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

    mux.HandleFunc("GET /{$}", home)  // Adds handle "/"-> func home()
    mux.HandleFunc("GET /snippet/view/{id}", snippetView)
    mux.HandleFunc("GET /snippet/create", snippetCreate)
    mux.HandleFunc("POST /snippet/create", snippetCreatePost)

    log.Print("starting server on :4000")

    err := http.ListenAndServe(":4000", mux)  // listens at port ("host:port")
    // if host not mentioned, goes an all network interfaces avaiable

    log.Fatal(err)
}
