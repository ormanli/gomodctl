package cmd

import (
	"github.com/ormanli/gomodctl/internal/module"
	"github.com/ormanli/gomodctl/internal/printer"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check [module name]",
	Short: "check local module for updates",
	Long:  `get list of local module and check them for updates`,
	RunE: func(cmd *cobra.Command, args []string) error {
		isJSON, err := cmd.Flags().GetBool("json")
		if err != nil {
			return err
		}

		path, err := cmd.Flags().GetString("path")
		if err != nil {
			return err
		}

		checkResults, err := module.Check(cmd.Context(), path)
		if err != nil {
			return err
		}

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
	rootCmd.AddCommand(checkCmd)
}
