package commands

import (
	"fmt"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/spf13/cobra"

	"github.com/BytemanD/easygo/pkg/fileutils"
	"github.com/BytemanD/easygo/pkg/global/logging"
)

var MDRender = &cobra.Command{
	Use:   "md-render <file path>",
	Short: "Markdown渲染",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fp := fileutils.FilePath{Path: args[0]}
		content, err := fp.ReadAll()
		if err != nil {
			logging.Fatal("read from file %s failed", fp.Path)
		}

		result := markdown.Render(content, 100, 4)
		fmt.Println(string(result))
	},
}
