package cmd

import (
	sa "github.com/kekeniker/spin-admin/cmd/service_account"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spin-admin",
	Short: "Spinnaker Admin Tools",
}

func init() {
	rootCmd.AddCommand(sa.NewServiceAccountCommand())
	rootCmd.AddCommand(newVersionCommand())
	rootCmd.AddCommand(newCompletionCmd())
}

func Execute() error {
	return rootCmd.Execute()
}
