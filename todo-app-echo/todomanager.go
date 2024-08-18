package main

import (
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	IsComplete bool   `json:"isDone"`
}

// required data to create new todo
type CreateTodoRequest struct {
	Title string `json:"title"`
}

type TodoManager struct {
	todos []Todo
	m     sync.Mutex // we want all our operations to be atomic
}

func NewTodoManager() TodoManager {
	return TodoManager{
		todos: make([]Todo, 0),
		m:     sync.Mutex{},
	}
}

func (tm *TodoManager) GetAll() []Todo {
	return tm.todos
}

func (tm *TodoManager) GetByID(ID string) (Todo, error) {
	var todo *Todo

	for _, todoItem := range tm.todos {
		if todoItem.ID == ID {
			todo = &todoItem

			break
		}
	}

	if todo == nil {
		return Todo{}, echo.ErrNotFound
	}

	return *todo, nil
}

func (tm *TodoManager) Create(createTodoRequest CreateTodoRequest) Todo {
	tm.m.Lock()
	defer tm.m.Unlock()

	newTodo := Todo{
		ID:         strconv.FormatInt(time.Now().UnixMilli(), 10),
		Title:      createTodoRequest.Title,
		IsComplete: false,
	}

	tm.todos = append(tm.todos, newTodo)

	return newTodo
}

func (tm *TodoManager) Complete(ID string) error {
	tm.m.Lock()
	defer tm.m.Unlock()

	// Find the todo with id
	var todo *Todo
	var index int = -1

	for indexKey, todoItem := range tm.todos {
		if todoItem.ID == ID {
			todo = &todoItem
			index = indexKey

			break
		}

		// if todoItem.ID != ID {
		// 	continue
		// }

		// todo = &todoItem
		// index = indexKey

		// break
	}

	if todo == nil {
		return echo.ErrNotFound
	}

	// Check todo is not already completed
	if todo.IsComplete {
		err := echo.ErrBadRequest
		err.Message = "todo is already complete"

		return err
	}

	// Update todo
	tm.todos[index].IsComplete = true

	return nil
}

func (tm *TodoManager) Remove(ID string) error {
	tm.m.Lock()
	defer tm.m.Unlock()

	index := -1

	for indexKey, todoItem := range tm.todos {
		if todoItem.ID == ID {
			index = indexKey

			break
		}
	}

	if index == -1 {
		return echo.ErrNotFound
	}

	tm.todos = append(tm.todos[:index], tm.todos[index+1:]...)

	return nil
}
