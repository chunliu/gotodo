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

func findTodo(id int) (int, Todo) {
	for i, t := range todoItems {
		if t.ID == id {
			return i, t
		}
	}

	return 0, Todo{}
}

func updateTodo(i int, t Todo) {
	todoItems[i].Name = t.Name
	todoItems[i].IsCompleted = t.IsCompleted
}

func deleteTodo(i int) {
	todoItems = append(todoItems[0:i], todoItems[i+1:]...) // delete a specific item from slice
}
