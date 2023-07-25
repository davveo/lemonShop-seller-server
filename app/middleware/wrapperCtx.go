package middleware

import (
	"context"
	"github.com/davveo/lemonShop-seller-server/app/consts"
	"github.com/gin-gonic/gin"
	"net/http"
)

const uuidKey = "uuid"

// WrapperCtx 可以在这个中间件里面获取前端header, cookie变量，传递到ctx里面
func WrapperCtx() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := ctx.Request.Header.Get(uuidKey)
		ctx.Request = AddToContext(ctx, uuidKey, uuid)
		ctx.Next()
	}
}

func AddToContext(c *gin.Context,
	key consts.ContextKey, value interface{}) *http.Request {
	return c.Request.WithContext(context.WithValue(c.Request.Context(), key, value))
}
