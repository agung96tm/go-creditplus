package main

import (
	"github.com/agung96tm/go-creditplus/internal/models"
	"github.com/agung96tm/go-creditplus/ui"
	"github.com/dustin/go-humanize"
	"github.com/justinas/nosurf"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
)

func priceComma(price float64) string {
	return humanize.Comma(int64(price))
}

var functions = template.FuncMap{
	"priceComma": priceComma,
}

type templateData struct {
	Form            any
	Limits          []*models.Limit
	Flash           string
	IsAuthenticated bool
	CSRFToken       string

	LoggedInUser *models.User
	User         *models.User
	Products     []*models.Product
	Product      *models.Product
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		Flash:           app.sessionManager.PopString(r.Context(), "flash"),
		IsAuthenticated: app.isAuthenticated(r),
		CSRFToken:       nosurf.Token(r),
	}
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
