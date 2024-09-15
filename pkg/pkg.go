package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/api-getway-todoapp-on-grapQl/config"
	"github.com/saidamir98/udevs_pkg/logger"
)

func ConnToDb(pgCfg config.PgConfig) (*pgx.Conn, error) {

	

	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		pgCfg.Username,
		pgCfg.Password,
		pgCfg.Host,
		pgCfg.Port,
		pgCfg.DatabaseName,
	)
	conn, err := pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		fmt.Println("error on Conn database", logger.Error(err))
		return nil, err
	}

	fmt.Println("successfuly conn with db ")
	return conn, nil
}
