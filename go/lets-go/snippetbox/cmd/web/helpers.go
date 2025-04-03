package main

import (
    "bytes"
    "errors"
	"fmt"
	"log/slog"
	"net/http"
    "runtime/debug"
    "time"

    "github.com/go-playground/form/v4" 
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack()) // stack trace
	)

	// should you use slog.Any(), slog.String(), etc?
	app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri), slog.String("trace", trace))
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Add common dynamic data
func (app *application) newTemplateData(r *http.Request) templateData {
    return templateData{
        CurrentYear: time.Now().Year(),
    }
}

// Function to render webpage from templateCache
func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	// retrieve required templateset and return it
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}

    // init buffer: catch runtime errors
    buf := new(bytes.Buffer)

	w.WriteHeader(status)

    // execute to buffer, not writer
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
	}

    // if executed without error, write to w
    buf.WriteTo(w)
}

func (app *application) decodePostForm(r *http.Request, dst any) error {
    // parse request body
    err := r.ParseForm()
    if err != nil {
        return err
    }

    // Use imported decoder (snippetCreateForm struct in handlers.go)
    err = app.formDecoder.Decode(dst, r.PostForm)
    if err != nil {
        // here we implement the special case
        // the decoder cannot have a nil form **and** this is a serious case
        // so we need a panic
        var invalidDecoderError *form.InvalidDecoderError

        if errors.As(err, &invalidDecoderError) {
            panic(err)
        }

        // all other cases, just have a bad request
        return err
    }

    // no error scenario
    return err
}
