package user

import (
	"context"
	"tfg/medical-center-service/pkg/config/errors"
	"tfg/medical-center-service/pkg/domain/models"
	"tfg/medical-center-service/pkg/domain/user"
	"tfg/medical-center-service/pkg/libs/crypto"
)

var (
	ErrorHashPassword = errors.Define("create_user.hash_password_error")
)

type CreateUser struct {
	repository user.Repository
}

func NewCreateUser(
	repository user.Repository,
) *CreateUser {
	return &CreateUser{
		repository: repository,
	}
}

type CreateUserDto struct {
	username string `json:"username"`
	password string `json:"password"`
}

type CreateUserResponse struct {
	ID string `json:id`
}

func (cu *CreateUser) Exec(ctx context.Context, payload *CreateUserDto) (*CreateUserResponse, error) {
	id := models.GenerateUUID()
	password, err := crypto.HashPassword(payload.password)
	if err != nil {
		return nil, err
	}
	err = cu.repository.Create(ctx, user.NewUser(id, payload.username, password))
	if err != nil {
		return nil, err
	}

	return &CreateUserResponse{ID: id.String()}, nil
}
