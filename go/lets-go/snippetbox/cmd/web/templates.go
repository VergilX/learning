package main

import (
    "html/template"
    "path/filepath"
    "time"

    "github.com/VergilX/learning/go/lets-go/snippetbox/internal/models"
)

type templateData struct {
    CurrentYear int
    Snippet     models.Snippet
    Snippets    []models.Snippet
    Form        any
}


// Custom functions in template
func humanDate(t time.Time) string {
    return t.Format("02 Jan 2006 at 15:04")
}


// store custom functions in single var
var functions = template.FuncMap{
    // key: used in template (frontend)
    "humanDate": humanDate,
}


func newTemplateCache() (map[string]*template.Template, error) {
    cache := map[string]*template.Template{}

    // returns paths satisfying the regex
    pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
    if err != nil {
        return nil, err
    }

    // loop through each file path
    for _, page := range pages {
        // extract final name "home.tmpl"
        name := filepath.Base(page)

        // register custom functions before calling parsing functions
        ts := template.New(name).Funcs(functions)

        // parse base template file into template set
        ts, err := ts.ParseFiles("./ui/html/base.tmpl")
        if err != nil {
            return nil, err
        }

        // call ParseGlob() **on this template set** to add any partials
        ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
        if err != nil {
            return nil, err
        }

        // call ParseFiles() **on this template set** to add the page template
        ts, err = ts.ParseFiles(page)
        if err != nil {
            return nil, err
        }

        // Add template set to map as cache
        cache[name] = ts
    }

    return cache, nil
}

