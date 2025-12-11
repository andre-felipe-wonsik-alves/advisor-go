package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adiciona uma nova tarefa",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("teste add")
		return nil
	},
}
