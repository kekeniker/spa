package service_account

import (
	"context"

	"github.com/kekeniker/spin-admin/pkg/client"
	"github.com/spf13/cobra"
)

func newServiceAccountCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		RunE:    serviceAccountCreateRun,
	}

	return cmd
}

func serviceAccountCreateRun(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	client, err := client.NewClient(ctx)
	if err != nil {
		return err
	}

	_ = client

	return nil
}
