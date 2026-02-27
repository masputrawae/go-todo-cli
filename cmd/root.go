/*
Copyright © 2026 Putra Jaya
*/
package cmd

import (
	"log"
	"os"

	"github.com/masputrawae/go-todo/pkg/model"
	"github.com/masputrawae/go-todo/utils"
	"github.com/spf13/cobra"
)

var Todos []model.Todo

var (
	task     string
	group    string
	priority string
)

var rootCmd = &cobra.Command{
	Use:   "go-todo",
	Short: "A brief description of your application",
	Long:  utils.Welcome(),
}

func Execute() {
	data, err := utils.LoadTodo("todos.json")
	if err != nil {
		log.Fatal(err)
	}

	Todos = data

	if err = rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
