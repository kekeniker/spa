package cmd

import (
	sa "github.com/kekeniker/spa/cmd/service_account"
	"github.com/kekeniker/spa/pkg/option"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spa",
	Short: "Spinnaker Admin Tools",
}

func init() {
	opt := &option.RootOption{}
	rootCmd.PersistentFlags().StringVarP(&opt.ConfigPath, "config", "c", "", "Debug output the operation")
	rootCmd.PersistentFlags().BoolVarP(&opt.Debug, "debug", "d", false, "Debug output the operation")

	rootCmd.AddCommand(sa.NewServiceAccountCommand(opt))
	rootCmd.AddCommand(newVersionCommand())
	rootCmd.AddCommand(newCompletionCmd())
}

func Execute() error {
	return rootCmd.Execute()
}
