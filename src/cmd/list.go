package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todas as tarefas.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("teste listar")
		return nil
	},
}
