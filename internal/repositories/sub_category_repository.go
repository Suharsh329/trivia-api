package repositories

import (
	"database/sql"
	"fmt"
	"time"
	"trivia/internal/models"
	"trivia/internal/utils"
)

type SubCategoryRepository struct {
	db *sql.DB
}

func NewSubCategoryRepository(db *sql.DB) *SubCategoryRepository {
	return &SubCategoryRepository{db: db}
}

func (r *SubCategoryRepository) GetAll(filters map[string]any) ([]models.SubCategory, error) {
	sql := `SELECT * FROM sub_categories`
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

	var subCategories []models.SubCategory
	for rows.Next() {
		var subCategory models.SubCategory
		if err = rows.Scan(
			&subCategory.ID,
			&subCategory.CategoryID,
			&subCategory.Name,
			&subCategory.ImageUrl,
			&subCategory.CreatedAt,
			&subCategory.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		subCategories = append(subCategories, subCategory)
	}

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return subCategories, nil
}

func (r *SubCategoryRepository) GetById(id int) (models.SubCategory, error) {
	stmt, err := r.db.Prepare(`SELECT * FROM sub_categories WHERE id = ?`)

	if err != nil {
		return models.SubCategory{}, fmt.Errorf("%v", err)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return models.SubCategory{}, fmt.Errorf("%v", err)
	}

	defer rows.Close()

	var subCategory models.SubCategory
	for rows.Next() {
		if err = rows.Scan(
			&subCategory.ID,
			&subCategory.CategoryID,
			&subCategory.Name,
			&subCategory.ImageUrl,
			&subCategory.CreatedAt,
			&subCategory.UpdatedAt); err != nil {
			return models.SubCategory{}, fmt.Errorf("%v", err)
		}
	}

	if err != nil {
		return models.SubCategory{}, fmt.Errorf("%v", err)
	}

	return subCategory, nil
}

func (r *SubCategoryRepository) Create(subCategory models.SubCategory) (int, error) {
	sql := `INSERT INTO sub_categories(name, category_id, image_url) VALUES(?, ?, ?) RETURNING id`

	stmt, err := r.db.Prepare(sql)

	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	var id int

	if err := stmt.QueryRow(subCategory.Name, subCategory.CategoryID, subCategory.ImageUrl).Scan(&id); err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	return id, nil
}

func (r *SubCategoryRepository) Update(subCategory models.SubCategory) error {
	sql := `UPDATE sub_categories SET name = ?, category_id = ?, image_url = ?, updated_at = ? WHERE id = ?`

	stmt, err := r.db.Prepare(sql)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	_, err = stmt.Exec(subCategory.Name, subCategory.CategoryID, subCategory.ImageUrl, time.Now().Format(time.RFC3339), subCategory.ID)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

func (r *SubCategoryRepository) Delete(id int) error {
	sql := `DELETE FROM sub_categories WHERE id = ?`

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
