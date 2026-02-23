package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/alecthomas/kong"
	"github.com/masputrawae/todo-cli/cmd"
	"github.com/masputrawae/todo-cli/internal/data"
	"github.com/masputrawae/todo-cli/internal/database"
	"github.com/masputrawae/todo-cli/internal/repo"
	"github.com/masputrawae/todo-cli/internal/service"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Printf("%T\n", ctx)

	db, err := database.SQLiteConnect("data/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repo.New(db)
	svc := service.New(repo, data.Statuses, data.Priorities)

	app := &cmd.App{
		Ctx: ctx,
		Svc: svc,
	}

	kctx := kong.Parse(&cmd.CLI)
	if err := kctx.Run(app); err != nil {
		log.Fatal(err)
	}
}
