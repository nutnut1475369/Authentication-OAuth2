package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID                   int                    `json:"id"        gorm:"primary_key"`
	UUID                 uuid.UUID              `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email                string                 `json:"email"`
	SocialId             string                 `json:"userId"`  
	LoginType			 int					`json:"loginType"`
	CreatedAt            time.Time              `json:"createdAt"`
	UpdatedAt            time.Time              `json:"updatedAt"`          
}
