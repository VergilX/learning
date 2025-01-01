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

        // slice with base, partials and pages
        files := []string{
            "./ui/html/base.tmpl",
            "./ui/html/partials/nav.tmpl",
            page,
        }

        // parse files into template set
        ts, err := template.ParseFiles(files...)
        if err != nil {
            return nil, err
        }

        // Add template set to map as cache
        cache[name] = ts
    }

    return cache, nil
}
