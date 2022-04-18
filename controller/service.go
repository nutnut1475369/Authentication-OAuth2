package controller

import (
	"encoding/json"
	"fmt"
	"googleauth/model"
	"googleauth/service/db"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	AccessToken string
}
type UserResponse struct {
	Id      string 
	Email   string 
}
var userRes UserResponse
func createUser(c *gin.Context, email string, socialId string, loginType int) model.User {
	fmt.Println(socialId)
	user := model.User{SocialId: socialId, LoginType: loginType}
	result := db.FromContext(c).DB.Find(&user)
	if result.RowsAffected > 0 {
		return user
	}
	newUser := model.User{Email: email, SocialId: socialId, LoginType: loginType}
	fmt.Println(newUser)
	db.FromContext(c).DB.Create(&newUser)
	return newUser
}
func PostTokenAfterGoogleSignIn(c *gin.Context) {
	var jsonAuth AuthRequest
	jsonAuth.AccessToken = c.GetHeader("Authorization")
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token="+jsonAuth.AccessToken)
	if  err != nil {
		c.JSON(http.StatusBadRequest,gin.H{})
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if  err != nil {
		c.JSON(http.StatusBadRequest,gin.H{})
	}
	fmt.Println(string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &userRes)
    if err != nil {
        fmt.Println(err)
    }
	if  err != nil {
        fmt.Println(err)
		return
	}
	createUser(c,userRes.Email,userRes.Id,1)
}