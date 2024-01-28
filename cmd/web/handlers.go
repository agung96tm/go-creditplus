package main

import (
	"errors"
	"github.com/agung96tm/go-creditplus/internal/models"
	"github.com/agung96tm/go-creditplus/internal/validator"
	"net/http"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/catalogs", http.StatusSeeOther)
}

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
	products, err := app.models.Product.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Products = products

	app.render(w, http.StatusOK, "catalog_list.tmpl", data)
}

func (app *application) catalogDetailHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.getParamId(r)
	if err != nil {
		app.notFound(w)
		return
	}

	product, err := app.models.Product.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoDataFound):
			app.notFound(w)
			return
		default:
			app.serverError(w, err)
			return
		}
	}

	data := app.newTemplateData(r)
	data.Product = product

	app.render(w, http.StatusOK, "catalog_detail.tmpl", data)
}

type BuyForm struct {
	LimitID             int    `form:"limit_id"`
	Name                string `form:"name"`
	Phone               string `form:"phone"`
	validator.Validator `form:"-"`
}

func (app *application) catalogBuyHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.getParamId(r)
	if err != nil {
		app.notFound(w)
		return
	}

	// product
	product, err := app.models.Product.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoDataFound):
			app.notFound(w)
		default:
			app.serverError(w, err)
		}
		return
	}

	// Logged-in user
	userID := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")
	user, err := app.models.User.GetById(userID)
	if err != nil {
		if errors.Is(err, models.ErrNoDataFound) {
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// your limit
	limits, err := app.models.Limit.GetLimitsByUserAndProduct(user, product)
	if err != nil {
		if errors.Is(err, models.ErrNoDataFound) {
			limits = make([]*models.Limit, 0)
		} else {
			app.serverError(w, err)
		}
	}

	data := app.newTemplateData(r)
	data.Product = product
	data.Limits = limits
	data.LoggedInUser = user
	data.Form = BuyForm{
		Name: user.FullName,
	}

	app.render(w, http.StatusOK, "catalog_buy.tmpl", data)
}

func (app *application) catalogBuyPostHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.getParamId(r)
	if err != nil {
		app.notFound(w)
		return
	}

	var form BuyForm
	err = app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Phone), "phone", "This field cannot be blank")
	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "catalog_buy.tmpl", data)
		return
	}

	// product
	product, err := app.models.Product.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoDataFound):
			app.notFound(w)
		default:
			app.serverError(w, err)
		}
		return
	}

	userID := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")
	user, err := app.models.User.GetById(userID)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoDataFound):
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		default:
			app.serverError(w, err)
		}
		return
	}

	limit, err := app.models.Limit.GetWithProductCalc(form.LimitID, product)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoDataFound):
			app.notFound(w)
		default:
			app.serverError(w, err)
		}
		return
	}

	trx, err := app.models.Transaction.Trx(
		form.Phone,
		user,
		limit,
		product,
	)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrLimitInsufficient):
			app.clientError(w, http.StatusBadRequest)
		case errors.Is(err, models.ErrNoMatched):
			panic("invalid limit and user")
		default:
			app.serverError(w, err)
		}
		return
	}

	_, err = app.models.Limit.ReduceLimit(limit, product.Price)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Transaction = trx
	data.Product = product
	data.Limit = limit

	app.render(w, http.StatusOK, "catalog_buy_done.tmpl", data)
}

func (app *application) dashboardHandler(w http.ResponseWriter, r *http.Request) {
	userID := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")
	user, err := app.models.User.GetById(userID)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoDataFound):
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		default:
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.LoggedInUser = user
	app.render(w, http.StatusOK, "dashboard.tmpl", data)
}
