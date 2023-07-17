package service

import (
	"api-blog/internal/entity"
	"api-blog/pkg/util"
	"context"
	"errors"
)

var (
	ErrWrongPassword = errors.New("password doesn't match")
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.User) error {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	err = m.Repository.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) Login(ctx context.Context, username, password string) error {
	compareHashedPassword, err := util.HashPassword(password)
	if err != nil {
		return err
	}

	user, err := m.Repository.Login(ctx, username)
	if err != nil {
		return err
	}

	if compareHashedPassword != user.Password {
		return ErrWrongPassword
	}

	return nil
}
