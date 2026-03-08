package store

import (
	"database/sql"
	"time"

	"github.com/erikdev/go-todo/models"
)

type TodoStore struct {
	db *sql.DB
}

func NewTodoStore(db *sql.DB) *TodoStore {
	return &TodoStore{db: db}
}

func (s *TodoStore) GetAll() ([]models.Todo, error) {
	rows, err := s.db.Query(
		`SELECT id, title, completed, created_at FROM todos ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var t models.Todo
		var completed int
		var createdAt string
		if err := rows.Scan(&t.ID, &t.Title, &completed, &createdAt); err != nil {
			return nil, err
		}
		t.Completed = completed == 1
		t.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		todos = append(todos, t)
	}
	return todos, rows.Err()
}

func (s *TodoStore) Create(title string) (int64, error) {
	result, err := s.db.Exec(`INSERT INTO todos (title) VALUES (?)`, title)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (s *TodoStore) Toggle(id int64) error {
	_, err := s.db.Exec(`UPDATE todos SET completed = NOT completed WHERE id = ?`, id)
	return err
}

func (s *TodoStore) Delete(id int64) error {
	_, err := s.db.Exec(`DELETE FROM todos WHERE id = ?`, id)
	return err
}
