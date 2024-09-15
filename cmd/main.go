package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jasurxaydarov/api-getway-todoapp-on-grapQl/config"
	"github.com/jasurxaydarov/api-getway-todoapp-on-grapQl/graphql/resolvers"
	"github.com/jasurxaydarov/api-getway-todoapp-on-grapQl/graphql/schema"
	db "github.com/jasurxaydarov/api-getway-todoapp-on-grapQl/pkg"
	"github.com/jasurxaydarov/api-getway-todoapp-on-grapQl/storage"
)

const defaultPort = "8080"

func main() {

	cfg := config.Load()

	conn, err := db.ConnToDb(cfg.PgConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	storage := storage.NewStorage(conn)

	srv := handler.NewDefaultServer(schema.NewExecutableSchema(schema.Config{Resolvers: &resolvers.Resolver{Storage: storage}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	go http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
