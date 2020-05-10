package main

// Todo represents the model of todo items.
type Todo struct {
	ID          int    `json:"id"`
	Key         int    `json:"key"`
	Name        string `json:"name"`
	IsCompleted bool   `json:"isCompleted"`
}

// TodoItems is an array of all todo items.
type TodoItems []Todo

// Config is the data model of gtd configurations.
type Config struct {
	BaseURL string `json:"baseURL"`
}
