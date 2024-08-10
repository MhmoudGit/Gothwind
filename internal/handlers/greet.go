package handlers

import (
	"go-comp/cmd/web/view/components"
	"go-comp/cmd/web/view/pages"
	"net/http"

	"github.com/labstack/echo/v4"
)

func BaseHandler(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Home())
}

func GreetPostHandler(c echo.Context) error {
	name := c.FormValue("name")
	return Render(c, http.StatusOK, components.GreetPost(name))
}
