package router

import (
	"fmt"
	"googleauth/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {
	gin.SetMode(gin.ReleaseMode)
	var r *gin.Engine
	if gin.Mode() == gin.ReleaseMode {
		r = gin.New()
	} else {
		r = gin.Default()
	}
	r.Use(middleware...)
	r.GET("/", func(c *gin.Context) {
		fmt.Print("hi")
	})
	r.POST("/auth/google", controller.PostTokenAfterGoogleSignIn)

	r.POST("/auth/facebook", controller.PostTokenAfterFacebookSignIn)

	r.GET("/logout/{provider}", func(c *gin.Context) {
		gothic.Logout(c.Writer, c.Request)
		c.Redirect(http.StatusTemporaryRedirect, "/")
	})
	r.Run()
	return r

}
