package cli

import "github.com/spf13/cobra"

type FlagInterface interface {
	RegistryTo(cmd *cobra.Command)
}

type StringFlag struct {
	Name      string
	ShortHand string
	Usage     string

	Default string
	value   *string
}

func (f *StringFlag) RegistryTo(cmd *cobra.Command) {
	if f.ShortHand != "" {
		f.value = cmd.Flags().StringP(f.Name, f.ShortHand, f.Default, f.Usage)
	}
	f.value = cmd.Flags().String(f.Name, f.Default, f.Usage)
}

func (f *StringFlag) Value() string {
	if *f.value != "" {
		return *f.value
	}
	return f.Default
}
