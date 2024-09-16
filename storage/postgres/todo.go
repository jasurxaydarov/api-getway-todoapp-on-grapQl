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
		fmt.Println("error on CreateTodo ", err)
		return nil, err

	}
	resp, err := t.GetTodo(ctx, &models.GetByID{ID: id.String()})

	if err != nil {
		fmt.Println("error on  Getodo", err)
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
		  is_completed,
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
		&resp.IsCompleted,
		&resp.CreatedAt,
	)

	if err != nil {
		fmt.Println("error on Getodo", logger.Error(err))
		return nil, err

	}

	return &resp, nil
}
func (t *todo) GetTodos(ctx context.Context, req *models.Gets) ([]*models.Todo, error) {

	var todo models.Todo
	var resp []*models.Todo
	offset := (req.Offset - 1) * req.Limit

	query := `
		SELECT
  		  todo_id,
		  task,
		  is_completed,
		  created_at
		FROM 
   			todos 
		WHERE
  			deleted_at IS NULL
		LIMIT $1 OFFSET $2
		`

	rows, err := t.conn.Query(ctx, query, req.Limit, offset)

	if err != nil {
		fmt.Println("error on GetTodos", logger.Error(err))
		return nil, err

	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&todo.ID,
			&todo.Task,
			&todo.IsCompleted,
			&todo.CreatedAt,
		)
		if err != nil {
			fmt.Println("error on GetTodos rows", logger.Error(err))
			return nil, err

		}

		resp = append(resp, &todo)

	}

	return resp, nil

}
func (t *todo) UpdateTodo(ctx context.Context, req *models.UpdateTodo) (*models.Todo, error) {
	query := `	
		UPDATE
			todos
		SET 
			task = $2,
			is_completed = $3,
			updated_at = current_timestamp
		WHERE 
			todo_id = $1;`

	_, err := t.conn.Exec(
		ctx,
		query,
		req.ID,
		req.Task,
		req.IsCompleted,
	)

	if err != nil {
		fmt.Println("error on UpdateTodo ", err)
		return nil, err

	}
	resp, err := t.GetTodo(ctx, &models.GetByID{ID: req.ID})

	if err != nil {
		fmt.Println("error on  Getodo", err)
		return nil, err

	}

	return resp, nil
}
func (t *todo) DeleteTodo(ctx context.Context, req *models.DeleteByID) (string, error) {
	query := `	
		DELETE FROM
			todos
		WHERE 
			todo_id = $1;`

	_, err := t.conn.Exec(
		ctx,
		query,
		req.ID,
	)

	if err != nil {
		fmt.Println("error on DeleteTodo ", err)
		return "", err

	}

	return "deleted succesfully", nil
}
