package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Task struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      string       `json:"status"`
	DueDate     time.Time    `json:"due_date"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"-"`
}

func (t *Task) CreateTask(db *sql.DB) error {
	query := "INSERT INTO tasks (title, description, status, due_date, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())"
	result, err := db.Exec(query, t.Title, t.Description, t.Status, t.DueDate)
	fmt.Println("errDDDB :", err)

	if err != nil {
		return err
	}
	t.ID, err = result.LastInsertId()
	return err
}

func (t *Task) UpdateTask(db *sql.DB) error {
	query := "UPDATE tasks SET title=?, description=?, status=?, due_date=?, updated_at=NOW() WHERE id=?"
	_, err := db.Exec(query, t.Title, t.Description, t.Status, t.DueDate, t.ID)
	return err
}

func DeleteTask(db *sql.DB, id int64) error {
	query := "UPDATE tasks SET deleted_at=NOW() WHERE id=?"
	_, err := db.Exec(query, id)
	return err
}

func GetTaskByID(db *sql.DB, id int64) (*Task, error) {
	query := "SELECT id, title, description, status, due_date, created_at, updated_at FROM tasks WHERE id=? AND deleted_at IS NULL"
	row := db.QueryRow(query, id)

	var task Task
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func GetAllTasks(db *sql.DB) ([]Task, error) {
	query := "SELECT id, title, description, status, due_date, created_at, updated_at FROM tasks WHERE deleted_at IS NULL ORDER BY due_date, status"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
