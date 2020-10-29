package controllers

import (
	"log"
	"testing"
	"time"
)

func TestCreateTodo(t *testing.T) {
	log.Printf("test")
	newTodo := Todo{
		ID:        "5",
		Title:     "test",
		Body:      "test",
		Completed: "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	CreateTodo(newTodo)
	if err != nil {
		t.Errorf("Insertion problem")
	}
}
