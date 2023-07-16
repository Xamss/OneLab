package pgrepo

import (
	"api-blog/internal/entity"
	"context"
	"fmt"
)

func (p *Postgres) CreateUser(ctx context.Context, u *entity.User) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                username, -- 1 
			                first_name, -- 2
			                last_name, -- 3
			                hashed_password -- 4
			                )
			VALUES ($1, $2, $3, $4)
			`, usersTable)

	_, err := p.Pool.Exec(ctx, query, u.Username, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Login(ctx context.Context, username, password string) (*entity.User, error) {
	query := fmt.Sprintf(`
		SELECT password FROM %s WHERE username = $1
	`, usersTable)
	var user entity.User

	err := p.Pool.QueryRow(ctx, query, username).Scan(&user.Password)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

// Need to fetch user data first - maybe use middleware and etc.
//func (p *Postgres) UpdateUser(ctx context.Context, u *entity.User) error{
//
//}
