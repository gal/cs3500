package controllers

import (
	"context"

	"github.com/gal/timber/models"
)

type UserController struct {
	Users models.UserStore
}

func NewUserController(userStore models.UserStore) *UserController {
	return &UserController{userStore}
}

func (userControl *UserController) Get(ctx context.Context, u *models.User) error {
	return userControl.Users.Get(ctx, u)
}

func (userControl *UserController) GetMany(ctx context.Context, userIds []int) ([]*models.User, error) {
	return userControl.Users.GetMany(ctx, userIds)
}

//update
func (userControl *UserController) Update(ctx context.Context, u *models.User) error {
	return userControl.Users.Update(ctx, u)
}
