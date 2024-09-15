package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/api-getway-todoapp-on-grapQl/models"
	"github.com/jasurxaydarov/api-getway-todoapp-on-grapQl/storage/repoi"
	"github.com/saidamir98/udevs_pkg/logger"
)

type todo struct {
	conn *pgx.Conn
}

func NewTodoRepo(conn *pgx.Conn) repoi.TodoRepoI {

	return &todo{conn: conn}
}

func (t *todo) CreateTodo(ctx context.Context, req *models.NewTodo) (*models.Todo, error) {
	id := uuid.New()

	query := `	
		INSERT INTO
			todos(
				todo_id,
				task
				
				
			)VALUES(
				$1,$2
			)`

	_, err := t.conn.Exec(
		ctx,
		query,
		id,
		req.Task,
	)

	if err != nil {
		fmt.Println("error on Createategory ", err)
		return nil, err

	}
	resp, err := t.GetTodo(ctx, &models.GetByID{ID: id.String()})

	if err != nil {
		fmt.Println("error on  GetCategory", err)
		return nil, err

	}

	return resp, nil

}
func (t *todo) GetTodo(ctx context.Context, req *models.GetByID) (*models.Todo, error) {
	var resp models.Todo

	query := `
		SELECT
  		  todo_id,
		  task,
		  created_at
		FROM 
   			todos 
		WHERE 
  			todo_id = $1;

	`

	err := t.conn.QueryRow(
		ctx,
		query,
		req.ID,
	).Scan(
		&resp.ID,
		&resp.Task,
		&resp.CreatedAt,
	
	)

	if err != nil {
		fmt.Println("error on GetComment", logger.Error(err))
		return nil, err

	}

	return &resp, nil
}
func (t *todo) GetTodos(ctx context.Context, req *models.Gets) (*[]models.Todo, error) {

	return nil, nil
}
func (t *todo) UpdateTodo(ctx context.Context, req *models.UpdateTodo) (*models.Todo, error) {

	return nil, nil
}
func (t *todo) DeleteTodo(ctx context.Context, req *models.DeleteByID) (string, error) {

	return "", nil
}
