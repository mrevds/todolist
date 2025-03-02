package service

import (
	"errors"
	"log"
	"time"
	"todo/internal/database"
	"todo/internal/model"
)

func CreateTask(title, description string) (model.Task, error) {
	if title == "" {
		return model.Task{}, errors.New("title cannot be empty")
	}

	task := model.Task{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	db := database.GetDB()
	query := `INSERT INTO tasks (title, description, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(query, task.Title, task.Description, task.CreatedAt, task.UpdatedAt).Scan(&task.ID)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func DeleteTask(id int) error {
	db := database.GetDB()
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting task from database: %v", err)
		return err
	}
	return nil
}

func UpdateTask(id int, title, description string) (model.Task, error) {
	if title == "" {
		return model.Task{}, errors.New("title cannot be empty")
	}

	task := model.Task{
		ID:          id,
		Title:       title,
		Description: description,
		UpdatedAt:   time.Now(),
	}

	db := database.GetDB()
	query := `UPDATE tasks SET title = $1, description = $2, updated_at = $3 WHERE id = $4 RETURNING created_at`
	err := db.QueryRow(query, task.Title, task.Description, task.UpdatedAt, task.ID).Scan(&task.CreatedAt)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func GetTask(id int) (model.Task, error) {
	task := model.Task{}
	db := database.GetDB()
	query := `SELECT id, title, description, created_at, updated_at FROM tasks WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func GetTasks() ([]model.Task, error) {
	tasks := []model.Task{}
	db := database.GetDB()
	query := `SELECT id, title, description, created_at, updated_at FROM tasks`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		task := model.Task{}
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}