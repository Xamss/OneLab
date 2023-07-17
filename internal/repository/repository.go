package repository

import (
	"api-blog/internal/entity"
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, u *entity.User) error
	Login(ctx context.Context, username string) (*entity.User, error)
	//UpdateUser(ctx context.Context, u *entity.User) error
	//DeleteUser(ctx context.Context, id int64) error
	//
	CreateArticle(ctx context.Context, a *entity.Article, u entity.User) error
	UpdateArticle(ctx context.Context, a *entity.Article) error
	DeleteArticle(ctx context.Context, id int64) error
	GetArticleByID(ctx context.Context, id int64) (*entity.Article, error)
	GetAllArticles(ctx context.Context) ([]entity.Article, error)
	GetArticlesByUserID(ctx context.Context, userID int64) ([]entity.Article, error)
	//
	GetCategories(ctx context.Context) ([]entity.Category, error)
	//GetCategoryByID(ctx context.Context) (entity.Category, error)

}
