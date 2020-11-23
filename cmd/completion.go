package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func newCompletionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "completion [bash|zsh|fish]",
		Short: "Generate a shell completion for spin-admin",
		Long: `To load completions:
	Bash:
	$ source <(spin-admin completion bash)
	# To load completions for each session, execute once:
	Linux:
	  $ spin-admin completion bash > /etc/bash_completion.d/spin-admin
	MacOS:
	  $ spin-admin completion bash > /usr/local/etc/bash_completion.d/spin-admin
	Zsh:
	# If shell completion is not already enabled in your environment you will need
	# to enable it.  You can execute the following once:
	$ echo "autoload -U compinit; compinit" >> ~/.zshrc
	# To load completions for each session, execute once:
	$ spin-admin completion zsh > "${fpath[1]}/_spin-admin"
	# You will need to start a new shell for this setup to take effect.
	Fish:
	$ spin-admin completion fish | source
	# To load completions for each session, execute once:
	$ spin-admin completion fish > ~/.config/fish/completions/spin-admin.fish
	`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "fish", "zsh"},
		Args:                  cobra.ExactValidArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			switch args[0] {
			case "zsh":
				err = cmd.Root().GenZshCompletion(os.Stdout)
			case "bash":
				err = cmd.Root().GenBashCompletion(os.Stdout)
			case "fish":
				err = cmd.Root().GenFishCompletion(os.Stdout, true)
			}
			if err != nil {
				return err
			}

			return nil
		},
	}
}
