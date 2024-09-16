package repoi

import (
	"context"

	"github.com/jasurxaydarov/api-getway-todoapp-on-grapQl/models"
)

type TodoRepoI interface {
	CreateTodo(ctx context.Context, req *models.NewTodo) (*models.Todo, error)
	GetTodo(ctx context.Context,req *models.GetByID)(*models.Todo,error)
	GetTodos(ctx context.Context,req *models.Gets)([]*models.Todo,error)
	UpdateTodo(ctx context.Context,req *models.UpdateTodo)(*models.Todo,error)
	DeleteTodo(ctx context.Context,req *models.DeleteByID)(string,error)
}
