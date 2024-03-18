package dto

import (
	"github.com/Max425/film-library.git/internal/common/constants"
	"github.com/Max425/film-library.git/internal/domain"
)

type SignInInput struct {
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignUpInputToDomainUser(signUp *SignUpInput) (*domain.User, error) {
	return domain.NewUser(0, signUp.Name, signUp.Mail, signUp.Password, "", constants.UserRole)
}
