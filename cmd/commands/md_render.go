package commands

import (
	"fmt"
	"io"
	"os"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/spf13/cobra"

	"github.com/BytemanD/easygo/pkg/fileutils"
	"github.com/BytemanD/easygo/pkg/terminal"
	"github.com/BytemanD/go-console/console"
)

var MDRender = &cobra.Command{
	Use:   "md-preview [file path]",
	Short: "Markdown预览",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		var err error
		var term *terminal.Terminal
		if len(args) != 0 {
			fp := fileutils.FilePath{Path: args[0]}
			content, err = fp.ReadAll()
			if err != nil {
				console.Fatal("read from file %s failed", fp.Path)
			}
			term = terminal.CurTerminal()
		} else {
			bytes, err := io.ReadAll(os.Stdin)
			if err != nil {
				console.Fatal("read from stdin failed")
			}
			content = string(bytes)
		}

		spredout, _ := cmd.Flags().GetBool("spredout")
		columns, leftPad := 100, 4
		if term != nil {
			if spredout {
				columns, leftPad = term.Columns, 0
			} else {
				columns, leftPad = term.Columns*2/3, term.Columns/6
			}
		} else {
			if spredout {
				leftPad = 0
			}
		}
		result := markdown.Render(content, columns+leftPad, leftPad)
		fmt.Println(string(result))
	},
}

func init() {
	MDRender.Flags().BoolP("spredout", "s", false, "铺满窗口, 默认不铺满, 宽度为窗口的2/3")
}
