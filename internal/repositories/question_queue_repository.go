package repositories

import (
	"database/sql"
	"fmt"
	"trivia/internal/models"
)

type QuestionQueueRepository struct {
	db *sql.DB
}

func NewQuestionQueueRepository(db *sql.DB) *QuestionQueueRepository {
	return &QuestionQueueRepository{db: db}
}

func (r *QuestionQueueRepository) FetchQueueByGameId(gameId, limit int64) ([]models.QuestionQueueResponse, error) {
	sql := "SELECT qq.id AS queue_id, qs.id AS question_id, qs.question_text, qs.correct_answer, qs.acceptable_answer, qs.difficulty_level, qs.image_url AS image, s.id AS sub_category_id, s.name, qq.game_id FROM question_queue qq INNER JOIN questions qs ON qq.question_id = qs.id INNER JOIN sub_categories s ON s.id = qs.sub_category_id WHERE qq.game_id = ? LIMIT ?"

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return []models.QuestionQueueResponse{}, fmt.Errorf("%v", err)
	}

	rows, err := stmt.Query(gameId, limit)
	if err != nil {
		return []models.QuestionQueueResponse{}, fmt.Errorf("%v", err)
	}
	defer rows.Close()

	var queue models.QuestionQueueResponse
	var queues []models.QuestionQueueResponse
	for rows.Next() {
		if err := rows.Scan(&queue.QueueID, &queue.QuestionID, &queue.QuestionText, &queue.CorrectAnswer, &queue.AcceptableAnswer, &queue.DifficultyLevel, &queue.SubCategoryID, &queue.SubCategoryName, &queue.GameID); err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		queues = append(queues, queue)
	}
	return queues, nil
}

func (r *QuestionQueueRepository) Add(gameId, questionId int) (int, error) {
	sql := "INSERT INTO question_queue (game_id, question_id) VALUES (?, ?) RETURNING id"
	stmt, err := r.db.Prepare(sql)

	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	var id int
	if err := stmt.QueryRow(gameId, questionId).Scan(&id); err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	return id, nil
}

func (r *QuestionQueueRepository) Delete(id int) error {
	sql := "DELETE FROM question_queue WHERE id = ?"

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
