package router

import (
	"encoding/json"
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
    r.POST("/auth/google", controller.PostTokenAfterGoogleSignIn )

    r.GET("/auth/:provider/callback", func(c *gin.Context) {
		q := c.Request.URL.Query()
		provider := c.Param("provider")
		q.Add("provider", provider)
		c.Request.URL.RawQuery = q.Encode()
		user,err := gothic.CompleteUserAuth(c.Writer,c.Request)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		res, err := json.Marshal(user)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		fmt.Printf(string(res))
    })

	r.GET("/logout/{provider}", func(c *gin.Context) {
		gothic.Logout(c.Writer,c.Request)
		c.Redirect(http.StatusTemporaryRedirect,"/")
	})
	r.Run()
	return r

}