package main

import (
	"github.com/agung96tm/go-creditplus/internal/models"
	"github.com/agung96tm/go-creditplus/ui"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

type templateData struct {
	Form      any
	Consumers []*models.Consumer
	Consumer  *models.Consumer
	Limits    []*models.Limit
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{}
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		patterns := []string{
			"html/base.tmpl",
			"html/partials/*.tmpl",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}

	return cache, nil
}
