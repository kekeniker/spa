package service_account

import "github.com/spf13/cobra"

func NewServiceAccountCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "service-account",
		Aliases: []string{"sa"},
	}

	cmd.AddCommand(newServiceAccountCreateCommand())

	return cmd
}
