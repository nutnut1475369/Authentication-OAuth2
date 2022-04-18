package db

import (
	"googleauth/model"

	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func (s *Service)Migrate() error {
	err := s.DB.AutoMigrate(
		&model.User{},
	) 
	if err !=nil {
		return err
	}
	return nil
}

func NewService(gdb *gorm.DB) (*Service,error){
	return &Service{DB: gdb}, nil
}