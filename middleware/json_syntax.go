package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func JSONSyntaxMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// decode request to byte array
		var body []byte
		if c.Request().Body != nil {
			body, _ = io.ReadAll(c.Request().Body)
		}

		// shoot down the request if there is a json syntax error
		if !json.Valid(body) && len(body) != 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "json syntax error",
			})
		}

		// encode the body back to io.ReadCloser
		c.Request().Body = io.NopCloser(bytes.NewBuffer(body))

		return next(c)
	}
}
