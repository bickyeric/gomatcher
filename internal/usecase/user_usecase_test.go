package usecase_test

import (
	"context"
	"testing"

	"github.com/bickyeric/gomatcher"
	"github.com/bickyeric/gomatcher/internal/usecase"
	"github.com/bickyeric/gomatcher/internal/usecase/mock"
	"go.uber.org/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockUserUsecase(ctrl)
	uc.EXPECT().CreateUser(context.Background(), gomatcher.StructIncludes{
		"Name": "John Doe",
	})

	uc.CreateUser(context.Background(), usecase.User{
		Name: "John Doe",
	})
}
