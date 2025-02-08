package cli

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Cli[T any] struct {
	Use           string
	Short         string
	ValidArgs     []cobra.PositionalArgs
	RegistryFlags func(fs *pflag.FlagSet) T
	Run           func(args []string, flags T) error
}

func NewCommand[T any](cli Cli[T]) *cobra.Command {
	if len(cli.ValidArgs) == 0 {
		cli.ValidArgs = []cobra.PositionalArgs{cobra.ExactArgs(0)}
	}

	var flags T

	cmd := &cobra.Command{
		Use: cli.Use,
		Args: func(cmd *cobra.Command, args []string) error {
			for _, validArg := range cli.ValidArgs {
				return validArg(cmd, args)
			}
			return nil
		},
		Short: cli.Short,
		Run: func(cmd *cobra.Command, args []string) {
			if cli.Run == nil {
				return
			}
			err := cli.Run(args, flags)
			if err != nil {
				os.Exit(1)
			}
		},
	}
	flags = cli.RegistryFlags(cmd.Flags())

	return cmd
}
