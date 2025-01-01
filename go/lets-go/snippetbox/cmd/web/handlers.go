package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/VergilX/learning/go/lets-go/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

    // add common dynamic data
    data := app.newTemplateData(r)
    data.Snippets = snippets

	// Using new render method (template cache)
	app.render(w, r, http.StatusOK, "home.tmpl", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// Retrieve wildcard values r.PathValue()
	// and error checking strconv
	id, err := strconv.Atoi(r.PathValue("id")) // voids non integers (Atoi)
	if err != nil || id < 1 {
		fmt.Println(err)
		http.NotFound(w, r)
		return
	}

	// Retrieve snippet from db using model
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}

		return
	}

    data := app.newTemplateData(r)
    data.Snippet = snippet

	// render page using helper render func
	app.render(w, r, http.StatusOK, "view.tmpl", data)
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
