package main

import (
    "flag"
    "log/slog"
    "net/http"
    "os"
)

func main() {
    // command line flags
    // flag function Args: 1-flagname; 2-defaultval; 3-desc
    addr := flag.String("addr", ":4000", "HTTP port address")

    // Right now addr has default value
    // Get the actual value, if present
    flag.Parse()
    
    // Structured logger
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

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

    logger.Info("starting server", slog.String("addr", *addr))

    err := http.ListenAndServe(*addr, mux)  // listens at port ("host:port")
    // if host not mentioned, goes an all network interfaces avaiable

    logger.Error(err.Error())
    os.Exit(1)
}
