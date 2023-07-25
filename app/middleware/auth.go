package middleware

import (
	"github.com/davveo/lemonShop-seller-server/conf"
	"github.com/davveo/lemonShop-seller-server/pkg/ojwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	whiteList = []string{""}
)

type HeaderParams struct {
	Authorization string `header:"Authorization"`
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Request.Header.Get("uuid")
		token := getToken(c)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "token鉴权未通过，请通过token授权接口重新获取token",
			})
			c.Abort()
		}

		if err := auth(c, token, uuid); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "token鉴权未通过，请通过token授权接口重新获取token",
			})
			c.Abort()
		}
		c.Next()
	}
}

func auth(c *gin.Context, token string, uuid string) error {
	claims, err := ojwt.NewJwt().ParseToken(token)
	if err != nil {
		return err
	}
	c.Request = AddToContext(c, "uid", claims.Data["uid"])
	c.Request = AddToContext(c, "username", claims.Data["username"])
	return nil
}

func getToken(c *gin.Context) string {
	var token string

	if conf.Conf.Env == "dev" || conf.Conf.Env == "test" {
		token = ""
	}
	// todo

	return token
}
