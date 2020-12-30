package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/golang-work/adventure/api/controllers/v1"
	"github.com/golang-work/adventure/api/middleware"
)

func Routers() *gin.Engine {
	var router = gin.Default()

	// router.Use(middleware.LoadTls())
	router.Use(middleware.RequestId(), middleware.Cors())

	guestGroup := router.Group("api")
	{
		guestGroup.POST("sms/send", v1.VcodeController().Send)

		guestGroup.POST("account/sign-up", v1.AccountController().SignUp)
		guestGroup.POST("account/sign-in", v1.AccountController().SignIn)

		guestGroup.GET("outside/login-server/list", v1.OutsideController().ListLoginServer)
	}

	privateGroup := router.Group("api", middleware.JWTAuth())
	{
		privateGroup.PUT("account/reset-password", v1.AccountController().ResetPassword)
		privateGroup.POST("account/retrieve-password", v1.AccountController().RetrievePassword)

		privateGroup.GET("sub-account", v1.SubAccountCrud().List)
		privateGroup.POST("sub-account", v1.SubAccountCrud().Store)
		privateGroup.DELETE("sub-account", v1.SubAccountCrud().Destroy)
		privateGroup.PUT("sub-account/recover", v1.SubAccountCrud().Recover)
	}
	return router
}
