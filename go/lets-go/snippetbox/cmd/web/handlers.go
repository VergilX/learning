package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
    "strings"
    "unicode/utf8"

	"github.com/VergilX/learning/go/lets-go/snippetbox/internal/models"
	"github.com/VergilX/learning/go/lets-go/snippetbox/internal/validator"
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

    // template tries to render .Form for previously values (in createPost)
    // but this is nil at first and causes error
    // this is to prevent this
    data.Form = snippetCreateForm {
        // can assign default values to fields
        Expires: 365,
    }

    app.render(w, r, http.StatusOK, "create.tmpl", data)
}

// struct for form data; also used for retrieving prev data
// in case of error
// (titled on purpose)
type snippetCreateForm struct {
    Title       string
    Content     string
    Expires     int
    validator.Validator
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

    // number so use Atoi func too
    // expires retrieved earlier (not in struct) for error detection
    expires, err := strconv.Atoi(r.PostForm.Get("expires"))
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    form := snippetCreateForm {
        Title:          r.PostForm.Get("title"),
        Content:        r.PostForm.Get("content"),
        Expires:        expires,
    }

    // Validate data in form variable using imported Validator

    // "title" should not be blank or >100chars
    form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
    form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")

    // "content" shouldn't be blank
    form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")

    // "expires" should be 1, 7 or 365
    form.CheckField(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

    if !form.Valid() {
        data := app.newTemplateData(r)
        data.Form = form
        app.render(w, r, http.StatusUnprocessableEntity, "create.tmpl", data)
        return
    }

    id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
    if err != nil {
        app.serverError(w, r, err)
        return
    }

    http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
