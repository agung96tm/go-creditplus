package main

import (
	"errors"
	"github.com/agung96tm/go-creditplus/internal/models"
	"github.com/agung96tm/go-creditplus/internal/validator"
	"net/http"
)

type LoginForm struct {
	NIK                 string `form:"nik"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = LoginForm{}
	app.render(w, http.StatusOK, "login.tmpl", data)
}

func (app *application) loginPostHandler(w http.ResponseWriter, r *http.Request) {
	var form LoginForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.NIK), "NIK", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Password), "Password", "This field cannot be blank")
	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "login.tmpl", data)
		return
	}

	id, err := app.models.User.Authenticate(form.NIK, form.Password)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrInvalidCredentials):
			form.AddNonFieldError("Email or Password is Incorrect")
			data := app.newTemplateData(r)
			data.Form = form
			app.render(w, http.StatusUnprocessableEntity, "login.tmpl", data)
		default:
			app.serverError(w, err)
		}
		return
	}

	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.sessionManager.Put(r.Context(), "authenticatedUserID", id)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (app *application) logoutHandler(w http.ResponseWriter, r *http.Request) {
	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Remove(r.Context(), "authenticatedUserID")
	app.sessionManager.Put(r.Context(), "flash", "You've been logged out successfully!")

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *application) catalogListHandler(w http.ResponseWriter, r *http.Request) {

}

func (app *application) catalogDetailHandler(w http.ResponseWriter, r *http.Request) {

}

func (app *application) catalogBuyHandler(w http.ResponseWriter, r *http.Request) {

}

func (app *application) catalogBuyPostHandler(w http.ResponseWriter, r *http.Request) {

}

func (app *application) dashboardHandler(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "dashboard.tmpl", data)
}
