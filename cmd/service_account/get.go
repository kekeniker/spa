package service_account

import (
	"context"
	"fmt"
	"os"

	"github.com/kekeniker/spa/pkg/client"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
)

type getOption struct {
	dryRun     bool
	namespace  string
	saOption   *saOption
	saName     string
	outputPath string
	username   string
}

func newServiceAccountGetCommand(sa *saOption) *cobra.Command {
	opts := &getOption{
		saOption: sa,
	}

	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"g"},
		RunE:    serviceAccountGetRun(opts),
	}

	cmd.PersistentFlags().BoolVarP(&opts.dryRun, "dryRun", "", defaultDryRun, "Dry run the operation")
	cmd.PersistentFlags().StringVarP(&opts.namespace, "namespace", "n", defaultNamespace, "Namespace to create the resources")
	cmd.PersistentFlags().StringVarP(&opts.saName, "service-account-name", "", defaultServiceAccountName, "Custom service account name")
	cmd.PersistentFlags().StringVarP(&opts.outputPath, "output", "o", "", "Path of the kube config output")
	cmd.PersistentFlags().StringVarP(&opts.username, "username", "u", defaultUsername, "AuthInfo username of the kubernetes context")
	return cmd
}

func serviceAccountGetRun(opt *getOption) func(cmd *cobra.Command, args []string) error {
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

		sa, err := client.GetServiceAccount(ctx, opt.saName, opt.namespace)
		if err != nil {
			return err
		}

		secret, err := client.GetSecret(ctx, sa.Secrets[0].Name, opt.namespace)
		if err != nil {
			return err
		}

		if opt.outputPath != "" {
			cfg, err := client.CreateKubeConfig(secret, opt.username)
			if err != nil {
				return err
			}

			if opt.outputPath == "-" {
				b, err := clientcmd.Write(*cfg)
				if err != nil {
					return err
				}

				fmt.Fprintf(os.Stdout, string(b))
				return nil
			}

			if err := clientcmd.WriteToFile(*cfg, opt.outputPath); err != nil {
				return err
			}

			return nil
		}

		// Print secret if no output path is specified
		fmt.Print(secret.Data["token"])
		return nil
	}
}
