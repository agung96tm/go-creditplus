package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	dynamic := alice.New()
	protected := dynamic.Append(app.requireAuthentication)

	router.Handler(http.MethodGet, "/catalogs", dynamic.ThenFunc(app.catalogListHandler))
	router.Handler(http.MethodGet, "/catalogs/:id", dynamic.ThenFunc(app.catalogDetailHandler))
	router.Handler(http.MethodGet, "/catalogs/:id/buy", protected.ThenFunc(app.catalogBuyHandler))
	router.Handler(http.MethodPost, "/catalogs/:id/buy", protected.ThenFunc(app.catalogBuyPostHandler))

	router.Handler(http.MethodGet, "/login", dynamic.ThenFunc(app.loginHandler))
	router.Handler(http.MethodPost, "/login", dynamic.ThenFunc(app.loginPostHandler))
	router.Handler(http.MethodGet, "/logout", protected.ThenFunc(app.logoutHandler))

	router.Handler(http.MethodGet, "/dashboard", protected.ThenFunc(app.dashboardHandler))

	standard := alice.New(app.recoverPanic, app.secureHeaders)
	return standard.Then(router)
}
