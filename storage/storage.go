package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/api-getway-todoapp-on-grapQl/storage/postgres"
	"github.com/jasurxaydarov/api-getway-todoapp-on-grapQl/storage/repoi"
)


type storage struct{

	todoRepoI repoi.TodoRepoI
}

type StorageI interface{

	GetTodoRepo()repoi.TodoRepoI
}


func NewStorage(db *pgx.Conn)StorageI{

	return &storage{
		todoRepoI: postgres.NewTodoRepo(db),
	}
}


func (s *storage)GetTodoRepo()repoi.TodoRepoI{

	return s.todoRepoI
}
