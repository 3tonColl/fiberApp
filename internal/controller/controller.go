package controller

import (
	"context"
	"github.com/3tonColl/fiberApp/fiberApp/model"
)

type UserController interface {
	GetUser(ctx context.Context, nickname, email string) (*model.User, error)
	InsertUser(ctx context.Context, User model.User) error
	UpdateUser(ctx context.Context, id int, UserUp model.UserUpd) (string, error)
	DeleteUser(ctx context.Context, id int) error
	GetUserPassword(ctx context.Context, id int) ([]byte, error)
	LoginUser(ctx context.Context, User *model.User) (string, error)
}

type userController struct {
	repo struct{}
}
