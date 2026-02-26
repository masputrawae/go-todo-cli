package cmd

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"

	"github.com/masputrawae/go-todo-cli/pkg/model"
)

type ListCmd struct {
	ID       *int    `short:"i" help:"Cari berdasarkan ID"`
	Status   *string `short:"s" help:"Cari Berdasarkan status"`
	Priority *string `short:"p" help:"Cari Berdasarkan Prioritas"`
	Category *string `short:"c" help:"Cari Berdasarkan Kategori"`
}

func (a *ListCmd) Run(app *App) error {
	var filters model.TodoFindType

	if a.ID != nil {
		filters.ID = a.ID
	}
	if a.Status != nil {
		filters.Status = a.Status
	}
	if a.Priority != nil {
		filters.Priority = a.Priority
	}
	if a.Category != nil {
		filters.Priority = a.Category
	}
	todos, err := app.Srv.GetTodo(&filters)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Task", "Status", "Priority", "Category", "Lastmod"})

	for _, v := range todos {
		lastmod := ""
		category := ""
		if v.Lastmod != nil {
			lastmod = v.Lastmod.Format("2006 Jan 01 15:02:01")
		}
		if v.Category != nil {
			category = *v.Category
		}
		table.Append(v.ID, v.Task, v.Status, v.Priority, category, lastmod)
	}
	table.Footer([]string{"Total", strconv.Itoa(len(todos))})

	table.Render()
	return nil
}
