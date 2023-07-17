package service

import (
	"api-blog/internal/entity"
	"context"
)

func (m *Manager) CreateArticle(ctx context.Context, a *entity.Article) error {
	//Fetch data later using Token or something
	var u entity.User
	err := m.Repository.CreateArticle(ctx, a, u)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) UpdateArticle(ctx context.Context, a *entity.Article) error {
	err := m.Repository.UpdateArticle(ctx, a)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) DeleteArticle(ctx context.Context, id int64) error {
	err := m.Repository.DeleteArticle(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) GetArticleByID(ctx context.Context, id int64) (*entity.Article, error) {
	article, err := m.Repository.GetArticleByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (m *Manager) GetAllArticles(ctx context.Context) ([]entity.Article, error) {
	articles, err := m.Repository.GetAllArticles(ctx)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (m *Manager) GetArticlesByUserID(ctx context.Context, userID int64) ([]entity.Article, error) {
	articles, err := m.Repository.GetArticlesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
