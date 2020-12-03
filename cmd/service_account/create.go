package service_account

import (
	"context"
	"fmt"
	"os"

	"github.com/kekeniker/spa/pkg/client"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

const (
	defaultDryRun             = false
	defaultServiceAccountName = "spinnaker"
	defaultRoleName           = "spinnaker-admin"
	defaultRoleBindingName    = "spinnaker-admin"
	defaultNamespace          = "kube-system"
)

type createOption struct {
	dryRun          bool
	isCluster       bool
	namespace       string
	roleName        string
	roleBindingName string
	saOption        *saOption
	saName          string
	outputPath      string
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
	cmd.PersistentFlags().StringVarP(&opts.roleBindingName, "role-binding-name", "", defaultRoleBindingName, "Custom role binding name")
	cmd.PersistentFlags().StringVarP(&opts.outputPath, "output", "o", "", "Path of the kube config output.")
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

		if opt.isCluster {
			role, err := client.CreateClusterRole(ctx, opt.roleName)
			if err != nil {
				return err
			}

			_, err = client.CreateClusterRoleBinding(ctx, sa.Name, role.Name, opt.roleBindingName, opt.namespace)
			if err != nil {
				return err
			}
		} else {
			role, err := client.CreateRole(ctx, opt.roleName, opt.namespace)
			if err != nil {
				return err
			}

			_, err = client.CreateRoleBinding(ctx, sa.Name, role.Name, opt.roleBindingName, opt.namespace)
			if err != nil {
				return err
			}
		}

		if opt.outputPath != "" {
			cfg, err := client.CreateKubeConfig(secret)
			if err != nil {
				return err
			}

			// TODO(KeisukeYamashita)
			// Create a DI container for mockable output

			b, err := yaml.Marshal(cfg)
			if err != nil {
				return err
			}

			if opt.outputPath == "-" {
				fmt.Print(string(b))
				return nil
			}

			f, err := os.OpenFile(opt.outputPath, os.O_RDONLY|os.O_CREATE|os.O_WRONLY, 0755)
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err := f.Write(b); err != nil {
				return err
			}

			return nil
		}

		return nil
	}
}
