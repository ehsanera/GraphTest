package customMiddleware

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/labstack/echo/v4"
)

func RateLimitMiddleware(limit *limiter.Limiter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			httpError := tollbooth.LimitByRequest(limit, c.Response(), c.Request())
			if httpError != nil {
				return c.JSON(httpError.StatusCode, map[string]string{
					"error": httpError.Message,
				})
			}
			return next(c)
		}
	}
}
