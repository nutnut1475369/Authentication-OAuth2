package services

import (
	"encoding/json"
	"fmt"
	"googleauth/model"
	"googleauth/service/db"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserResponse struct {
	Id      string 
	Email   string 
}

type MyCustomClaims struct {
	UUID string `json:"uuid"`
	jwt.StandardClaims
}

var userRes *UserResponse
func createUser(c *gin.Context, email string, socialId string, loginType int) model.User {
	user := model.User{SocialId: socialId, LoginType: loginType}
	result := db.FromContext(c).DB.Where("social_id = ? AND login_type = ?",socialId,loginType).Find(&user)
	fmt.Print(result.RowsAffected)
	if result.RowsAffected > 0 {
		return user
	}
	fmt.Print(user)
	newUser := model.User{Email: email, SocialId: socialId, LoginType: loginType}
	db.FromContext(c).DB.Create(&newUser)
	return newUser
}
func getFacebookProfile(accessToken string) (*UserResponse, error){

	client := &http.Client{
        Timeout: time.Second * 10,
    }
	url := "https://graph.facebook.com/me?fields=email"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
        return nil,err
    }
	req.Header.Add("Authorization",accessToken)
	resp,err := client.Do(req)
	if  err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if  err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &userRes)
	if  err != nil {
		return nil, err
	}
	return userRes,err
}

func getGoogleProfile(accessToken string) (*UserResponse, error){
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token="+accessToken)
	if  err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if  err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &userRes)
	if  err != nil {
		return nil, err
	}
	return userRes,err
}

func CreateGoogleUser(c *gin.Context, accessToken string) (string, error){
	userRes, err := getGoogleProfile(accessToken)
	if  err != nil {
		return "", err
	}
	user := createUser(c,userRes.Email,userRes.Id,1)
	claim := MyCustomClaims{
		user.UUID.String(),
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "Server",
		},
	}
	// fmt.Print(claim)
	token,err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(os.Getenv("SIGNING_KEY")))
	return token, err
}
func CreateFacebookUser(c *gin.Context, accessToken string) (string, error){
	userRes, err := getFacebookProfile(accessToken)
	if  err != nil {
		return "", err
	}
	user := createUser(c,userRes.Email,userRes.Id,0)
	claim := MyCustomClaims{
		user.UUID.String(),
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "Server",
		},
	}
	// fmt.Print(claim)
	token,err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(os.Getenv("SIGNING_KEY")))
	return token, err
}