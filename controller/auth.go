package controller

import (
	"net/http"

	"googleauth/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthRequest struct {
	AccessToken string `json:"accessToken" validate:"required"`
}

type AuthHeader struct {
	Authorization string `header:"authorization"`
}

type AuthResponse struct {
	Status      string `json:"status"`
	AccessToken string `json:"accessToken"`
}
var validate *validator.Validate

func PostTokenAfterGoogleSignIn(c *gin.Context) {
	var json AuthRequest
	json.AccessToken = c.GetHeader("Authorization")

	validate = validator.New()
	err := validate.Struct(json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.CreateGoogleUser(c, json.AccessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, AuthResponse{AccessToken: token, Status: "success"})
}
func PostTokenAfterFacebookSignIn(c *gin.Context) {
	var json AuthRequest
	json.AccessToken = c.GetHeader("Authorization")
	validate = validator.New()
	err := validate.Struct(json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.CreateFacebookUser(c, json.AccessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, AuthResponse{AccessToken: token, Status: "success"})
}