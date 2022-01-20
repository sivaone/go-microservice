package service

import (
	"context"
	"fmt"
	"log"

	"go-microservice/config"
	"go-microservice/ent"
)

type UserOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewUserOps(ctx context.Context) *UserOps {
	return &UserOps{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (r *UserOps) UsersGetAll() ([]*ent.User, error) {

	users, err := r.client.User.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserOps) CreateUser(user ent.User) (*ent.User, error) {

	createdUser, err := r.client.User.Create().
		SetName(user.Name).
		SetAge(user.Age).
		Save(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	log.Println("user is created: ", createdUser)
	return createdUser, nil
}
