/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/masputrawae/go-todo/model"
	"github.com/masputrawae/go-todo/todo"
	"github.com/spf13/cobra"
)

var (
	id       int
	task     string
	group    string
	done     bool
	priority string
)

// init todos
var Todos = todo.New([]model.Todo{}, "todos.json")

var rootCmd = &cobra.Command{
	Use:   "go-todo",
	Short: "A brief description of your application",
}

func Execute() {
	if err := Todos.Load(); err != nil {
		log.Fatal(err)
	}
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
