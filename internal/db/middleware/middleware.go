package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupMiddleware(e *echo.Echo) {
	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.CORS(),
		middleware.RequestID(),
	)
}
