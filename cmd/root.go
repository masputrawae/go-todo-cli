package cmd

import (
	"context"
	"fmt"

	"github.com/masputrawae/todo-cli/internal/repo"
	"github.com/masputrawae/todo-cli/internal/service"
)

type App struct {
	Ctx context.Context
	Svc service.Manage
}

var CLI struct {
	Add    AddCmd    `cmd:"" help:"Add new todo"`
	Update UpdateCmd `cmd:"" help:"Update todo"`
	Delete DeleteCmd `cmd:"" help:"Delete todo"`
	List   ListCmd   `cmd:"" help:"List todos"`
}

// ===== ADD CMD
type AddCmd struct {
	Task     *string `arg:"" required:"" help:"Task description"`
	Priority *string `short:"p" help:"Priority"`
	Status   *string `short:"s" help:"Status"`
	Category *string `short:"c" help:"Category"`
}

func (a *AddCmd) Run(app *App) error {
	if err := app.Svc.Create(app.Ctx, &repo.CreateInputType{
		Task:     a.Task,
		Status:   a.Status,
		Priority: a.Priority,
		Category: a.Category,
	}); err != nil {
		return err
	}
	return nil
}

// ===== UPDATE CMD
type UpdateCmd struct {
	ID       *int    `arg:"" required:""`
	Task     *string `short:"t" help:"Task description"`
	Priority *string `short:"p" help:"Priority"`
	Status   *string `short:"s" help:"Status"`
	Category *string `short:"c" help:"Category"`
}

func (a *UpdateCmd) Run(app *App) error {
	if err := app.Svc.Update(app.Ctx, &repo.UpdateInputType{
		ID:       a.ID,
		Task:     a.Task,
		Status:   a.Status,
		Priority: a.Priority,
		Category: a.Category,
	}); err != nil {
		return err
	}
	return nil
}

// ===== DELETE CMD
type DeleteCmd struct {
	ID *int `arg:"" required:""`
}

func (a *DeleteCmd) Run(app *App) error {
	return app.Svc.Delete(app.Ctx, &repo.DeleteInputType{ID: a.ID})
}

// ===== LIST CMD
type ListCmd struct {
	ID       *int    `short:"i"`
	Priority *string `short:"p" help:"Priority"`
	Status   *string `short:"s" help:"Status"`
	Category *string `short:"c" help:"Category"`
}

func (a *ListCmd) Run(app *App) error {
	todos, err := app.Svc.GetTodos(app.Ctx, &repo.FindInputType{
		ID:       a.ID,
		Status:   a.Status,
		Priority: a.Priority,
		Category: a.Category,
	})
	if err != nil {
		return err
	}

	for i, v := range todos.Todos {
		fmt.Printf("%d: ID: %d\nStatus: %s\nPriority: %s\nCategory: %s\nTask: %s\nLastmod: %s\n\n",
			i+1,
			v.ID,
			v.Status,
			v.Priority,
			v.Category,
			v.Task,
			v.Lastmod.Time,
		)
	}
	return nil
}
