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
