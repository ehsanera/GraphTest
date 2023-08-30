package customMiddleware

import (
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"strings"
)

type limitedReaderCloser struct {
	io.Reader
}

func (l *limitedReaderCloser) Close() error {
	if closer, ok := l.Reader.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

func MinBodySizeMiddleware(minSize int64) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
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
