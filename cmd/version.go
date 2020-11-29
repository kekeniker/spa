package cmd

import (
	"fmt"

	"github.com/kekeniker/spa/pkg/version"
	"github.com/spf13/cobra"
)

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.String())
		},
	}
}
