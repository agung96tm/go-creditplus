package main

import (
	assert "github.com/agung96tm/go-creditplus/internal/asset"
	"net/http"
	"net/url"
	"testing"
)

func doLogin(t *testing.T, ts *testServer) {
	_, _, body := ts.Get(t, "/login")
	actualToken := extractCSRFToken(t, body)

	form := url.Values{}
	form.Add("nik", "123456789")
	form.Add("password", "pa$$word")
	form.Add("csrf_token", actualToken)
	ts.PostForm(t, "/login", form)
}

func TestHomeView(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, _ := ts.Get(t, "/")
	assert.Equal(t, code, http.StatusSeeOther)
}

func TestCatalogListView(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.Get(t, "/catalogs")
	assert.Equal(t, code, 200)
	assert.StringContains(t, body, "Product1")
}

func TestCatalogDetailView(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.Get(t, "/catalogs/1")
	assert.Equal(t, code, 200)
	assert.StringContains(t, body, "Product1")
}

func TestDashboardView(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	doLogin(t, ts)
	code, _, body := ts.Get(t, "/dashboard")
	assert.Equal(t, code, http.StatusOK)
	assert.StringContains(t, body, "123456789")
}
