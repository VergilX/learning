package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strconv"
)


func (app *application) home(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Server", "Go")

    // files slice (base template should be *first*)
    files := []string{
        "./ui/html/base.tmpl",
        "./ui/html/pages/home.tmpl",
        "./ui/html/partials/nav.tmpl",
    }

    // ts: Template set
    // takes files slice
    // ... is used to pass slice elements as args
    ts, err := template.ParseFiles(files...)
    if err != nil {
        app.serverError(w, r, err)  // helpers.go
        return
    }

    // Add template to http body
    // second arg: name of template ( {{define "base"}} )
    // last arg: dynamic data (nil for now)
    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        app.serverError(w, r, err)  // helpers.go
    }
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
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

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Display a form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
    // Write http header
    title := "0 snail"
    content := "0 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
    expires := 7

    id, err := app.snippets.Insert(title, content, expires)
    if err != nil {
        app.serverError(w, r, err)
        return
    }

        http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
