package pgrepo

import (
	"api-blog/internal/entity"
	"context"
	"fmt"
)

func (p *Postgres) CreateArticle(ctx context.Context, a *entity.Article, u entity.User) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (title, description, user_id, categories)
		VALUES($1, $2, $3, $4)
	`, articleTable)

	var categories_id []int64
	for _, c := range a.Categories {
		categories_id = append(categories_id, c.ID)
	}
	_, err := p.Pool.Exec(ctx, query, a.Title, a.Description, u.ID, categories_id)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) UpdateArticle(ctx context.Context, a *entity.Article) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET title = $1, description = $2, categories = $3
		WHERE id = $4
	`)

	var categories_id []int64
	for _, c := range a.Categories {
		categories_id = append(categories_id, c.ID)
	}
	_, err := p.Pool.Exec(ctx, query, a.Title, a.Description, categories_id, a.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) DeleteArticle(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s 
		WHERE  id = $1
	`, articleTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) GetArticleByID(ctx context.Context, id int64) (*entity.Article, error) {
	query := fmt.Sprintf(`
		SELECT title, description, categories FROM %s
		WHERE id = $1
	`, articleTable)

	var article entity.Article
	err := p.Pool.QueryRow(ctx, query, id).Scan(&article.Title, &article.Description, &article.Categories)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (p *Postgres) GetAllArticles(ctx context.Context) ([]entity.Article, error) {
	query := fmt.Sprintf(`
		SELECT title, description, categories
		FROM %s
	`, articleTable)

	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var articles []entity.Article
	for rows.Next() {
		var article entity.Article
		err := rows.Scan(
			&article.Title,
			&article.Description,
			&article.Categories,
		)
		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (p *Postgres) GetArticlesByUserID(ctx context.Context, userID int64) ([]entity.Article, error) {
	query := fmt.Sprintf(`
		SELECT title, description, categories FROM %s
		WHERE user_id = $1
	`)

	rows, err := p.Pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	articles := []entity.Article{}
	for rows.Next() {
		var article entity.Article
		err = rows.Scan(
			&article.Title,
			&article.Description,
			&article.Categories,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}
