package cli

import (
	"context"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "email-validator",
	Short: "API for email validation",
}

func Execute(ctx context.Context) error {
	initCommands(rootCmd)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		return err
	}

	return nil
}

func initCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(
		NewServe(),
	)
}
