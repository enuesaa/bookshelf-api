package apptest

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/router/middleware"
	"github.com/labstack/echo/v4"
)

func New(t *testing.T, repos repository.Repos) AppTest {
	return AppTest{
		t: t,
		repos: repos,
	}
}

type AppTest struct {
	t *testing.T
	repos repository.Repos
}

func (a *AppTest) Run(method string, path string, handler echo.HandlerFunc, body io.Reader) (*httptest.ResponseRecorder, error) {
	if err := a.repos.Startup(); err != nil {
		return nil, err
	}

	app := echo.New()
	app.Use(middleware.BindCtx(a.repos))
	app.Use(middleware.HandleData)
	app.Use(middleware.HandleError)

	app.Any(path, handler)

	req := httptest.NewRequest(method, path, nil)
	rec := httptest.NewRecorder()

	app.ServeHTTP(rec, req)

	return rec, nil
}

func (a *AppTest) Get(path string, handler echo.HandlerFunc) (*httptest.ResponseRecorder, error) {
	return a.Run("GET", path, handler, nil)
}

func (a *AppTest) Post(path string, handler echo.HandlerFunc, body string) (*httptest.ResponseRecorder, error) {
	return a.Run("POST", path, handler, strings.NewReader(body))
}
