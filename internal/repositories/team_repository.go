package repositories

import (
	"database/sql"
	"fmt"
	"trivia/internal/models"
	"trivia/internal/utils"
)

type TeamRepository struct {
	db *sql.DB
}

func NewTeamRepository(db *sql.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

func (r *TeamRepository) Get(filters map[string]any) ([]models.Team, error) {
	sql := `SELECT * FROM teams`

	if len(filters) > 0 {
		exactMatch, ok := filters["exactMatch"].(bool)
		if !ok {
			exactMatch = false
		}
		sql += utils.CreateQueryFilters(filters, exactMatch)
	}

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return nil, fmt.Errorf("TeamRepository@Get | Error: %v", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("TeamRepository@Get | Error: %v", err)
	}
	defer rows.Close()

	teams := []models.Team{}
	for rows.Next() {
		team := models.Team{}
		if err := rows.Scan(&team.ID, &team.GameID, &team.Name, &team.Score, &team.CreatedAt, &team.UpdatedAt); err != nil {
			return nil, fmt.Errorf("TeamRepository@Get | Error: %v", err)
		}
		teams = append(teams, team)
	}
	return teams, nil
}

func (r *TeamRepository) Create(team models.Team) (int, error) {
	sql := `INSERT INTO teams (name, game_id) VALUES (?, ?) RETURNING id`

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return 0, fmt.Errorf("TeamRepository@Create | Error: %v", err)
	}

	var id int
	if err := stmt.QueryRow(team.Name, team.GameID).Scan(&id); err != nil {
		return 0, fmt.Errorf("TeamRepository@Create | Error: %v", err)
	}
	return id, nil
}

func (r *TeamRepository) Update(team models.Team) error {
	sql := `UPDATE teams SET name = ? WHERE id = ?`
	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return fmt.Errorf("TeamRepository@Update | Error: %v", err)
	}

	if _, err := stmt.Exec(team.Name, team.ID); err != nil {
		return fmt.Errorf("TeamRepository@Update | Error: %v", err)
	}
	return nil
}

func (r *TeamRepository) Delete(id int) error {
	sql := `DELETE FROM teams WHERE id = ?`
	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return fmt.Errorf("TeamRepository@Delete | Prepare Error: %v", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("TeamRepository@Delete | Exec Error: %v", err)
	}
	return nil
}
