package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-work/adventure/api/protocol"
	"github.com/golang-work/adventure/foundation"
	"github.com/golang-work/adventure/support"
	"strconv"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			protocol.Response(ctx).Abort("unauthenticated").Json()
			ctx.Abort()
			return
		}
		jwt := foundation.NewJWT()
		claims, err := jwt.ParseToken(token)
		if err != nil {
			e := support.Abort{}
			if err == foundation.TokenExpired {
				e = support.Throw("tokenExpired")
			} else {
				e = support.Throw("tokenInvalid", "reason", err.Error())
			}
			protocol.Response(ctx).Abort(e).Json()
			ctx.Abort()
			return
		}
		if claims.ExpiresAt-support.Now().Unix() < claims.RefreshTtl {
			claims.ExpiresAt = support.Now().Unix() + support.Config["auth"].GetInt64("jwt.ttl")
			newToken, _ := jwt.CreateToken(*claims)
			newClaims, _ := jwt.ParseToken(newToken)
			ctx.Header("At-New-Token", newToken)
			ctx.Header("At-New-Token-ExpiresAt", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
