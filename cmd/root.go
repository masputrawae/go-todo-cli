/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/masputrawae/todo-cli/internal/model"
	"github.com/masputrawae/todo-cli/internal/repo"
	"github.com/masputrawae/todo-cli/internal/service"
	"github.com/masputrawae/todo-cli/internal/utils"
	"github.com/spf13/cobra"
)

var (
	Todos    []model.Todo
	FilePath = "todos.json"
	Cfg      = utils.LoadConfig("configs/config.yaml")
	Repo     = repo.NewTodoRepo(Todos, FilePath)
	Svc      = service.NewTodoService(Repo, Cfg.Statuses, Cfg.Priorities)
)

var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := Svc.GetAll()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, v := range todos {
			fmt.Printf("[%d] ", v.ID)
			if v.Priority != nil {
				fmt.Printf("(%s) ", *v.Priority)
			}
			fmt.Printf("%s ", v.Task)
			if v.Priority != nil {
				fmt.Printf("%s ", *v.Status)
			}
			if v.Project != nil {
				fmt.Printf("| %s ", *v.Project)
			}
		}
	},
}

func Execute() {
	if err := Svc.Load(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
