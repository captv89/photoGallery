package web

import (
	"embed"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Embed the static files
//
//go:embed static
var content embed.FS

func StartServer() {
	e := echo.New()

	// Load the static files from the embed.FS
	e.StaticFS("/assets", content)

	// Load middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Load the routes
	setupRoutes(e)

	// Start the server
	e.Logger.Fatal(e.Start("localhost:3000"))
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}
