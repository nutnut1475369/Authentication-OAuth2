package context

import (
	"context"
	"googleauth/service/db"
	"googleauth/service/jwt"

	"github.com/gin-gonic/gin"
)
func SetDb(s *db.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		db.ToContext(c, s)
		ctx := context.WithValue(c.Request.Context(), db.Key, s)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
func SetJwt(s *jwt.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt.ToContext(c, s)
		ctx := context.WithValue(c.Request.Context(), jwt.Key, s)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
