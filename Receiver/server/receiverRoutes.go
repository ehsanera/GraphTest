package server

import (
	"Receiver/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http"
	"strings"
)

func receiverRoutes(e *echo.Echo) {
	e.Use(middleware.BodyLimit("8K"))
	e.Use(minBodySizeMiddleware(50))

	e.POST("/send", postHandler)
}

func postHandler(c echo.Context) error {
	send := &models.SendRequest{}
	if err := c.Bind(send); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, send)
}

type limitedReaderCloser struct {
	io.Reader
}

func (l *limitedReaderCloser) Close() error {
	if closer, ok := l.Reader.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

func minBodySizeMiddleware(minSize int64) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Read the first `minSize` bytes from the request body
			body := c.Request().Body
			buf := make([]byte, minSize)
			n, err := io.ReadFull(body, buf)
			if err != nil {
				return c.String(http.StatusBadRequest, "Request body is too small")
			}

			limitedBody := &limitedReaderCloser{
				Reader: io.MultiReader(io.NopCloser(strings.NewReader(string(buf[:n]))), body),
			}

			c.Request().Body = limitedBody
			return next(c)
		}
	}
}
