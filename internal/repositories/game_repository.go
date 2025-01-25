package repositories

import (
	"database/sql"
	"fmt"
	"trivia/internal/models"
	"trivia/internal/utils"
)

type GameRepository struct {
	db *sql.DB
}

func NewGameRepository(db *sql.DB) *GameRepository {
	return &GameRepository{db: db}
}

func (r *GameRepository) Get(filters map[string]any) ([]models.Game, error) {
	sql := `SELECT * FROM games`

	if len(filters) > 0 {
		exactMatch, ok := filters["exactMatch"].(bool)
		if !ok {
			exactMatch = false
		}
		sql += utils.CreateQueryFilters(filters, exactMatch)
	}

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return nil, fmt.Errorf("GameRepository@Get | Error: %v", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("GameRepository@Get | Error: %v", err)
	}
	defer rows.Close()

	games := []models.Game{}
	for rows.Next() {
		game := models.Game{}
		if err := rows.Scan(&game.ID, &game.Name, &game.CreatedAt, &game.UpdatedAt); err != nil {
			return nil, fmt.Errorf("GameRepository@Get | Error: %v", err)
		}
		games = append(games, game)
	}
	return games, nil
}

func (r *GameRepository) Create(game models.Game) (models.Game, error) {
	sql := `INSERT INTO games (name) VALUES (?) RETURNING id, name`

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return models.Game{}, fmt.Errorf("GameRepository@Create | Error: %v", err)
	}

	newGame := models.Game{}
	if err := stmt.QueryRow(game.Name).Scan(&newGame.ID, &newGame.Name); err != nil {
		return models.Game{}, fmt.Errorf("GameRepository@Create | Error: %v", err)
	}
	return newGame, nil
}

func (r *GameRepository) Update(game models.Game) error {
	sql := `UPDATE games SET name = ? WHERE id = ?`
	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return fmt.Errorf("GameRepository@Update | Error: %v", err)
	}

	if _, err := stmt.Exec(game.Name, game.ID); err != nil {
		return fmt.Errorf("GameRepository@Update | Error: %v", err)
	}
	return nil
}

func (r *GameRepository) Delete(id int) error {
	sql := `DELETE FROM games WHERE id = ?`
	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return fmt.Errorf("GameRepository@Delete | Prepare Error: %v", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("GameRepository@Delete | Exec Error: %v", err)
	}
	return nil
}
