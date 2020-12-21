package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func LoadTls() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		})
		err := middleware.Process(ctx.Writer, ctx.Request)
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.Next()
	}
}
