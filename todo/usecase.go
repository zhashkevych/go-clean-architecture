package todo

type UseCase interface {
	RegisterUser()
	CreateTodo()
	GetUserTodos()
	GetTodoByID()
	DeleteTodo()
}