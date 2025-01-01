package main

import (
    "net/http"

    "github.com/justinas/alice"
)

// update to return Handler. Doesn't break the code as Handler
// also satisfies interface of ListenAndServe() in main.go
func (app *application) routes() http.Handler {
    mux := http.NewServeMux()  // one to many (multiplexer by regex)

    // Create a fileserver for static files
    fileserver := http.FileServer(http.Dir("./ui/static/"))

    // Handler for fileserver
    // all URLs that start with /static/
    // Arg taken is stripped of /static using http.StripPrefix
    mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

    mux.HandleFunc("GET /{$}", app.home)  // Adds handle "/"-> func home()
    mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
    mux.HandleFunc("GET /snippet/create", app.snippetCreate)
    mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

    // chain middleware using imported package
    standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

    // returns the servemux
    return standard.Then(mux)
}
