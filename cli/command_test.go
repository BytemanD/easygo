package cli

import (
	"testing"

	"github.com/spf13/pflag"
)

func TestCli(t *testing.T) {
	type DemoFlags struct {
		Name *string
	}
	cmd := NewCommand(Cli[DemoFlags]{
		RegistryFlags: func(fs *pflag.FlagSet) DemoFlags {
			return DemoFlags{
				Name: fs.String("name", "", ""),
			}
		},
		Run: func(args []string, flags DemoFlags) error {
			println("name =", *flags.Name)
			return nil
		},
	})
	cmd.SetArgs([]string{"--name", "foo"})
	if err := cmd.Execute(); err != nil {
		t.Errorf("run command failed: %s", err)
	}
}
