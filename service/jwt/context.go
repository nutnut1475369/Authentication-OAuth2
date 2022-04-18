package jwt

import (
	"context"
)

type Setter interface {
	Set(string, interface{})
}

const Key = "jwt"

func FromContext(c context.Context) *Service {
	return c.Value(Key).(*Service)
}

func ToContext(c Setter, s *Service) {
	c.Set(Key, s)
}
