package main

var currentID int
var todoItems TodoItems

func init() {
	RepoAddTodoItem(Todo{Name: "Todo Item 1", IsCompleted: false})
}

func RepoAddTodoItem(item Todo) Todo {
	item.ID = currentID
	todoItems = append(todoItems, item)
	currentID++
	return item
}

func RepoFindTodo(id int) Todo {
	for _, t := range todoItems {
		if t.ID == id {
			return t
		}
	}

	return Todo{}
}
