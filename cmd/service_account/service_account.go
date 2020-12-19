package service_account

import (
	"github.com/kekeniker/spa/pkg/option"
	"github.com/spf13/cobra"
)

type saOption struct {
	rootOpt *option.RootOption
}

func NewServiceAccountCommand(rootOpt *option.RootOption) *cobra.Command {
	opt := &saOption{
		rootOpt: rootOpt,
	}

	cmd := &cobra.Command{
		Use:     "service-account",
		Aliases: []string{"sa"},
	}

	cmd.AddCommand(newServiceAccountGetCommand(opt))
	cmd.AddCommand(newServiceAccountCreateCommand(opt))
	return cmd
}
