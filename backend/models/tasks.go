package models

import "github.com/google/uuid"

type Task struct {
	ID 			string `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (t *Task) GenerateID(){
	t.ID = uuid.NewString()
} 