package repositories

import (
	"database/sql"
	"fmt"
	"time"
	"trivia/internal/models"
	"trivia/internal/utils"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll(filters map[string]any) ([]models.Category, error) {
	sql := `SELECT * FROM categories`
	stmt, err := r.db.Prepare(sql)

	if len(filters) > 0 {
		exactMatch, ok := filters["exactMatch"].(bool)
		if !ok {
			exactMatch = false
		}
		sql += utils.CreateQueryFilters(filters, exactMatch)
	}

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err = rows.Scan(
			&category.ID,
			&category.Name,
			&category.ImageUrl,
			&category.CreatedAt,
			&category.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		categories = append(categories, category)
	}

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return categories, nil
}

func (r *CategoryRepository) GetById(id int) (models.Category, error) {
	stmt, err := r.db.Prepare(`SELECT * FROM categories WHERE id = ?`)

	if err != nil {
		return models.Category{}, fmt.Errorf("%v", err)
	}

	rows, err := stmt.Query(id)

	if err != nil {
		return models.Category{}, fmt.Errorf("%v", err)
	}

	defer rows.Close()

	var category models.Category
	for rows.Next() {
		if err = rows.Scan(
			&category.ID,
			&category.Name,
			&category.ImageUrl,
			&category.CreatedAt,
			&category.UpdatedAt); err != nil {
			return models.Category{}, fmt.Errorf("%v", err)
		}
	}

	if err != nil {
		return models.Category{}, fmt.Errorf("%v", err)
	}

	return category, nil
}

func (r *CategoryRepository) Create(category models.Category) (int, error) {
	sql := `INSERT INTO categories(name, image_url) VALUES(?, ?) RETURNING id`

	stmt, err := r.db.Prepare(sql)

	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	var id int

	if err := stmt.QueryRow(category.Name, category.ImageUrl).Scan(&id); err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	return id, nil
}

func (r *CategoryRepository) Update(category models.Category) error {
	sql := `UPDATE categories SET name = ?, image_url = ?, updated_at = ? WHERE id = ?`

	stmt, err := r.db.Prepare(sql)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	_, err = stmt.Exec(category.Name, category.ImageUrl, time.Now().Format(time.RFC3339), category.ID)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

func (r *CategoryRepository) Delete(id int) error {
	sql := `DELETE FROM categories WHERE id = ?`

	stmt, err := r.db.Prepare(sql)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	_, err = stmt.Exec(id)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}
