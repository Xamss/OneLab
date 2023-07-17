package service

import (
	"api-blog/internal/entity"
	"context"
)

func (m *Manager) GetCategories(ctx context.Context) ([]entity.Category, error) {
	categories, err := m.Repository.GetCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
