package cmd

import (
	"fmt"

	"github.com/andre-felipe-wonsik-alves/advisor-go/internal/task"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adiciona uma nova tarefa",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("teste add")
		store := task.NewJSONStore("data/tasks.json")
		tasks, _ := store.Load()

		return nil
	},
}
