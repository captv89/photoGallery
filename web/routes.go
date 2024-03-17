package web

import (
	"github.com/labstack/echo/v4"
)

func setupRoutes(e *echo.Echo) {
	// Home page
	e.GET("/", home)
	
	// Image Modal
	e.GET("/image/:id", imgModal)

	
}
