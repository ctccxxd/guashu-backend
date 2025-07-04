package middleware

import (
	"github.com/labstack/echo/v4"
	"jihulab.com/guashu/gs-server/lib/jwt"
	"jihulab.com/guashu/gs-server/util/xrsp"
	"net/http"
	"time"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("token")
		if token == "" {
			token = c.QueryParam("token")
		}

		tokenData, err := jwt.ParseToken(token)
		if err != nil || tokenData.UserId == 0 || tokenData.ExpireAt < time.Now().Unix() {
			return c.JSON(http.StatusUnauthorized, xrsp.ErrorText("invalid token"))
		}

		c.Set("user_id", tokenData.UserId)

		return next(c)
	}
}
