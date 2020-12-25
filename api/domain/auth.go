package domain

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang-work/adventure/foundation"
	"github.com/golang-work/adventure/models"
	"github.com/golang-work/adventure/support"
	"go.uber.org/zap"
)

type auth struct {
	ctx   *gin.Context
	token string
}

func Auth(ctx *gin.Context) *auth {
	return &auth{
		ctx: ctx,
	}
}

func (d *auth) User() *foundation.Claims {
	claims, _ := d.ctx.Get("claims")
	return claims.(*foundation.Claims)
}

func (d *auth) MakeToken(account *models.Account) (string, int64, error) {
	j := &foundation.JWT{SigningKey: []byte(support.Config["auth"].GetString("jwt.signingKey"))}
	claims := foundation.Claims{
		ID:         account.ID,
		Username:   account.Username,
		RefreshTtl: support.Config["auth"].GetInt64("jwt.refreshTtl"),
		StandardClaims: jwt.StandardClaims{
			NotBefore: support.Now().Unix() - 1000,
			ExpiresAt: support.Now().Unix() + support.Config["auth"].GetInt64("jwt.ttl"),
			Issuer:    account.Username,
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		support.Log.Error(err.Error(), zap.Any("err", err))
		return "", 0, support.Throw("jwtMakeTokenFail")
	}

	return token, claims.StandardClaims.ExpiresAt, nil
}
