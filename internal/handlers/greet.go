package handlers

import (
	"gothwind/cmd/web/view/components"
	"gothwind/cmd/web/view/pages"
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
