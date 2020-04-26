package main

// TodoRepo is the repository for all todo items.
type TodoRepo struct {
	CurrentID int
	Items     TodoItems
}

var todoRepo TodoRepo

func init() {
	todoRepo.CurrentID = 1 // Id start with 1
	todoRepo.Add(Todo{Name: "Create Web API with Go", IsCompleted: true})
	todoRepo.Add(Todo{Name: "Build front end", IsCompleted: false})
}

// Add function adds a new todo item to the repo.
func (tr *TodoRepo) Add(item Todo) Todo {
	item.ID = tr.CurrentID
	item.Key = tr.CurrentID
	tr.Items = append(tr.Items, item)
	tr.CurrentID++
	return item
}

// Find function finds a specific item according to the value of id.
func (tr *TodoRepo) Find(id int) (int, Todo) {
	for i, t := range tr.Items {
		if t.ID == id {
			return i, t
		}
	}

	return 0, Todo{}
}

// Update function updates an item according to its index in the slice.
func (tr *TodoRepo) Update(i int, t Todo) {
	tr.Items[i].Name = t.Name
	tr.Items[i].IsCompleted = t.IsCompleted
}

// Delete function deletes an item based on its index.
func (tr *TodoRepo) Delete(i int) {
	tr.Items = append(tr.Items[0:i], tr.Items[i+1:]...) // delete a specific item from slice
}
