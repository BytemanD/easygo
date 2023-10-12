package commands

import (
	"fmt"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/spf13/cobra"

	"github.com/BytemanD/easygo/pkg/fileutils"
	"github.com/BytemanD/easygo/pkg/global/logging"
	"github.com/BytemanD/easygo/pkg/terminal"
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

		// width, _ := cmd.Flags().GetInt("width")
		// leftAlign, _ := cmd.Flags().GetInt("left-align")
		spredout, _ := cmd.Flags().GetBool("spredout")

		columns, leftPad := 100, 4
		terminal := terminal.CurTerminal()
		if terminal != nil {
			fmt.Println(terminal.Columns)
			if spredout {
				columns, leftPad = terminal.Columns, 0
			} else {
				columns, leftPad = terminal.Columns*2/3, terminal.Columns/6
			}
		}
		result := markdown.Render(content, columns+leftPad, leftPad)
		fmt.Println(string(result))
	},
}

func init() {
	MDRender.Flags().BoolP("spredout", "s", false, "铺满窗口, 默认不铺满, 宽度为窗口的2/3")
}
