package usecase

type UseCase interface {
	RegisterUser()
	CreateTodo()
	GetUserTodos()
	GetTodoByID()
	DeleteTodo()
}