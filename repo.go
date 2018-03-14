package main

var currentID int
var todoItems TodoItems

func init() {
	addTodoItem(Todo{Name: "Todo Item 1", IsCompleted: false})
}

func addTodoItem(item Todo) Todo {
	item.ID = currentID
	todoItems = append(todoItems, item)
	currentID++
	return item
}

func findTodo(id int) Todo {
	for _, t := range todoItems {
		if t.ID == id {
			return t
		}
	}

	return Todo{}
}
