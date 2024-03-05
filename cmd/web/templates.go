package main

import (
	"fmt"
	"html/template"
	"path/filepath"

	"snippetbox.charygarry.net/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	fmt.Printf("pages = %v\n", pages)

	for _, page := range pages {
		name := filepath.Base(page)

		fmt.Printf("page = %v\n", page)

		// Parse the base template file into a template set.
		ts, err := template.ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		// Call ParseGlob() *on this template set* to add any partials.
		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		// Call ParseFiles() *on this template set* to add the page template.
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		fmt.Printf("ts = %v\n", ts)
		cache[name] = ts
	}

	return cache, nil
}
