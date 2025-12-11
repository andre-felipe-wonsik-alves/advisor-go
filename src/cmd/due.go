package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dueCmd = &cobra.Command{
	Use:   "due",
	Short: "Verifica tarefas com vencimentos próximos/vencidos.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("teste due")
		return nil
	},
}
