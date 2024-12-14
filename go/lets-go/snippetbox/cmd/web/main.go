package main

import (
    "flag"
    "log"
    "net/http"
)

func main() {
    // command line flags
    // flag function Args: 1-flagname; 2-defaultval; 3-desc
    addr := flag.String("addr", ":4000", "HTTP port address")

    // Right now addr has default value
    // Get the actual value, if present
    flag.Parse()

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

    log.Printf("starting server on %s", *addr)

    err := http.ListenAndServe(*addr, mux)  // listens at port ("host:port")
    // if host not mentioned, goes an all network interfaces avaiable

    log.Fatal(err)
}
