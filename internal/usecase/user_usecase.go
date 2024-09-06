package usecase

import "context"

type User struct {
	ID   int64
	Name string
}

//go:generate mockgen -destination=mock/user_service.go -package=mock github.com/bickyeric/gomatcher/internal/usecase UserUsecase
type UserUsecase interface {
	CreateUser(context.Context, *User) error
}
