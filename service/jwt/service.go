package jwt

import (
	"golang.org/x/oauth2/jwt"
)

type Service struct {
	JWT  *jwt.Config
}

func NewService(jwt *jwt.Config) (*Service,error){
	return &Service{JWT: jwt}, nil
}