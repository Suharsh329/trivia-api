package repositories

import (
	"database/sql"
	"fmt"
	"trivia/internal/models"
	"trivia/internal/utils"
)

type PlayerRepository struct {
	db *sql.DB
}

func NewPlayerRepository(db *sql.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

func (r *PlayerRepository) Get(filters map[string]any) ([]models.Player, error) {
	sql := `SELECT * FROM players`

	if len(filters) > 0 {
		exactMatch, ok := filters["exactMatch"].(bool)
		if !ok {
			exactMatch = false
		}
		sql += utils.CreateQueryFilters(filters, exactMatch)
	}

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return nil, fmt.Errorf("PlayerRepository@Get | Error: %v", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("PlayerRepository@Get | Error: %v", err)
	}
	defer rows.Close()

	players := []models.Player{}
	for rows.Next() {
		player := models.Player{}
		if err := rows.Scan(&player.ID, &player.GameID, &player.Name, &player.Score, &player.CreatedAt, &player.UpdatedAt); err != nil {
			return nil, fmt.Errorf("PlayerRepository@Get | Error: %v", err)
		}
		players = append(players, player)
	}
	return players, nil
}

func (r *PlayerRepository) Create(player models.Player) (int, error) {
	sql := `INSERT INTO players (name, game_id) VALUES (?, ?) RETURNING id`

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return 0, fmt.Errorf("PlayerRepository@Create | Error: %v", err)
	}

	var id int
	if err := stmt.QueryRow(player.Name, player.GameID).Scan(&id); err != nil {
		return 0, fmt.Errorf("PlayerRepository@Create | Error: %v", err)
	}
	return id, nil
}

func (r *PlayerRepository) Update(player models.Player) error {
	sql := `UPDATE players SET name = ? WHERE id = ?`
	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return fmt.Errorf("PlayerRepository@Update | Error: %v", err)
	}

	if _, err := stmt.Exec(player.Name, player.ID); err != nil {
		return fmt.Errorf("PlayerRepository@Update | Error: %v", err)
	}
	return nil
}

func (r *PlayerRepository) Delete(id int) error {
	sql := `DELETE FROM players WHERE id = ?`
	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return fmt.Errorf("PlayerRepository@Delete | Prepare Error: %v", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("PlayerRepository@Delete | Exec Error: %v", err)
	}
	return nil
}
