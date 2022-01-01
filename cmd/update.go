package cmd

import (
	"github.com/ormanli/gomodctl/internal/module"
	"github.com/ormanli/gomodctl/internal/printer"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update project dependencies",
	Long:  `update project dependencies to minor versions`,
	RunE: func(cmd *cobra.Command, args []string) error {
		isJSON, err := cmd.Flags().GetBool("json")
		if err != nil {
			return err
		}

		path, err := cmd.Flags().GetString("path")
		if err != nil {
			return err
		}

		checkResults, err := module.Update(cmd.Context(), path)
		if err != nil {
			return err
		}

		cmd.Println("Your dependencies updated to latest minor and go.mod.backup created")

		rp := NewResultPrinter(checkResults)
		if isJSON {
			printer.PrintJSON(rp)
		} else {
			printer.PrintTable(rp)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
