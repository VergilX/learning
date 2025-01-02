package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
    "strings"
    "unicode/utf8"

	"github.com/VergilX/learning/go/lets-go/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
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
    data := app.newTemplateData(r)

    app.render(w, r, http.StatusOK, "create.tmpl", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
    // Limit size of request body to 4096 bytes
    r.Body = http.MaxBytesReader(w, r.Body, 4096)

    // parses request body
    err := r.ParseForm()
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    // extract form data
    title := r.PostForm.Get("title")
    content := r.PostForm.Get("content")

    // number so use Atoi func too
    expires, err := strconv.Atoi(r.PostForm.Get("expires"))
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    // Validate data
    fieldErrors := make(map[string]string)

    // "title" should not be blank or >100chars
    if strings.TrimSpace(title) == "" {
        fieldErrors["title"] = "This field cannot be blank"
    } else if utf8.RuneCountInString(title) > 100 {
        fieldErrors["title"] = "This field cannot be more than 100 chars long"
    }

    // "content" cannot be blank
    if strings.TrimSpace(content) == "" {
        fieldErrors["content"] = "This field cannot be blank"
    }

    // "expires" should be either 1, 7 or 365
    if expires != 1 && expires != 7 && expires != 365 {
        fieldErrors["expires"] = "This field must be equal to 1, 7 or 365"
    }

    if len(fieldErrors) > 0 {
        fmt.Fprint(w, fieldErrors)
        return
    }

    id, err := app.snippets.Insert(title, content, expires)
    if err != nil {
        app.serverError(w, r, err)
        return
    }

    http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
