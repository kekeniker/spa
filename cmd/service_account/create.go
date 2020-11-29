package service_account

import (
	"context"

	"github.com/kekeniker/spa/pkg/client"
	"github.com/spf13/cobra"
)

const (
	defaultDryRun             = false
	defaultServiceAccountName = "spinnaker"
	defaultRoleName           = "spinnaker-admin"
	defaultNamespace          = "kube-system"
)

type createOption struct {
	dryRun    bool
	namespace string
	roleName  string
	saOption  *saOption
	saName    string
	isCluster bool
}

func newServiceAccountCreateCommand(sa *saOption) *cobra.Command {
	opts := &createOption{
		saOption: sa,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		RunE:    serviceAccountCreateRun(opts),
	}

	cmd.PersistentFlags().BoolVarP(&opts.isCluster, "cluster", "", false, "Cluster level role and bindings or not")
	cmd.PersistentFlags().BoolVarP(&opts.dryRun, "dryRun", "", defaultDryRun, "Dry run the operation")
	cmd.PersistentFlags().StringVarP(&opts.namespace, "namespace", "n", defaultNamespace, "Namespace to create the resources")
	cmd.PersistentFlags().StringVarP(&opts.saName, "service-account-name", "", defaultServiceAccountName, "Custom service account name")
	cmd.PersistentFlags().StringVarP(&opts.roleName, "role-name", "", defaultRoleName, "Custom role name")
	return cmd
}

func serviceAccountCreateRun(opt *createOption) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		opts := []client.ClientOption{}
		if opt.dryRun {
			opts = append(opts, client.WithDryRun())
		}

		client, err := client.NewClient(ctx, opts...)
		if err != nil {
			return err
		}

		sa, secret, err := client.CreateServiceAccount(ctx, opt.saName, opt.namespace)
		if err != nil {
			return err
		}

		role, err := client.CreateRole(ctx, sa.Name, opt.roleName)
		if err != nil {
			return err
		}

		rolebinding, err := client.CreateRoleBinding(ctx, sa.Name, role.Name, opt.namespace)
		if err != nil {
			return err
		}

		_ = secret
		_ = rolebinding

		return nil
	}
}
