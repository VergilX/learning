package main

import (
    "html/template"
    "path/filepath"

    "github.com/VergilX/learning/go/lets-go/snippetbox/internal/models"
)

type templateData struct {
    Snippet models.Snippet
    Snippets []models.Snippet
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

        // parse base template file into template set
        ts, err := template.ParseFiles("./ui/html/base.tmpl")
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
