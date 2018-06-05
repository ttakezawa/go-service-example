package main

import (
	"log"
	"os"

	"github.com/facebookgo/inject"
	"github.com/ttakezawa/go-service-example/app/server"
	"github.com/ttakezawa/go-service-example/infrastructure/postgresql"
	"github.com/ttakezawa/go-service-example/usecase"
)

func main() {
	if err := run(); err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
}

func run() error {
	var server server.App
	db, err := postgresql.Open(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)
	if err != nil {
		return err
	}

	var g inject.Graph
	err = g.Provide(
		&inject.Object{Value: &server},
		&inject.Object{Value: db},
		&inject.Object{Value: &usecase.UserUsecase{}},
		&inject.Object{Value: &postgresql.UserRepository{}},
	)
	if err != nil {
		return err
	}

	if err = g.Populate(); err != nil {
		return err
	}

	return server.Run()
}
