package pgrepo

import (
	"api-blog/internal/entity"
	"context"
	"fmt"
)

func (p *Postgres) GetCategories(ctx context.Context) ([]entity.Category, error) {
	query := fmt.Sprintf(`
		SELECT id, name
		FROM %s
	`, categoryTable)

	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		var category entity.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
