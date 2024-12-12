package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Server", "Go")

    w.Write([]byte("Hello from SnippetBox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
    // Retrieve wildcard values r.PathValue()
    // and error checking strconv
    id, err := strconv.Atoi(r.PathValue("id"))  // voids non integers (Atoi)
    if err != nil || id < 1 {
        fmt.Println(err)
        http.NotFound(w, r)
        return
    }

    /*
    // Sprintf allows interpolation with message (old way)
    msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
    w.Write([]byte(msg))
    */

    // use a function instead of primitive Write method
    // uses interface io.Writer as argument
    // so can use standard library functions like fmt
    fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Display a form for creating a new snippet..."))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
    // Write http header
    w.WriteHeader(http.StatusCreated)

    w.Write([]byte("Save a new snippet"))
}

func main() 
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
