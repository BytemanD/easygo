package commands

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/BytemanD/easygo/pkg/fileutils"
	"github.com/BytemanD/go-console/console"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	// "github.com/BytemanD/easygo/pkg/table"
	"github.com/spf13/cobra"
)

var TABLE_STYLES = map[string]table.Style{
	"default": table.StyleDefault,
	"light":   table.StyleLight,
	"rounded": table.StyleRounded,
	"bold":    table.StyleBold,
}
var TABLE_COLORS = map[string]text.Color{
	"cyan":   text.FgCyan,
	"blue":   text.FgBlue,
	"yellow": text.FgYellow,
	"red":    text.FgRed,
	"green":  text.FgGreen,
}

func getRow(columns []string, color ...text.Color) table.Row {
	renderFunc := fmt.Sprintf
	if len(color) > 0 {
		renderFunc = color[0].Sprintf
	}
	tableRow := table.Row{}
	for _, column := range columns {
		tableRow = append(tableRow, renderFunc(column))
	}
	return tableRow
}

var CSVRender = &cobra.Command{
	Use:   "csv-render [file path]",
	Short: "CSV格式渲染表格",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		noHeader, _ := cmd.Flags().GetBool("no-header")
		style, _ := cmd.Flags().GetString("style")
		color, _ := cmd.Flags().GetString("color")

		tableWriter := table.NewWriter()

		if _, ok := TABLE_STYLES[style]; ok {
			tableWriter.SetStyle(TABLE_STYLES[style])
		} else {
			fmt.Println("invalid style:", style)
			os.Exit(1)
		}
		if color != "" {
			if _, ok := TABLE_COLORS[color]; ok {
				tableWriter.Style().Color.Border = text.Colors{TABLE_COLORS[color]}
				tableWriter.Style().Color.Separator = text.Colors{TABLE_COLORS[color]}
			} else {
				fmt.Println("invalid color:", color)
				os.Exit(1)
			}
		}

		var (
			content string
			err     error
		)
		if len(args) != 0 {
			fp := fileutils.FilePath{Path: args[0]}
			content, err = fp.ReadAll()
			if err != nil {
				console.Fatal("read from file %s failed", fp.Path)
			}
		} else {
			bytes, err := io.ReadAll(os.Stdin)
			if err != nil {
				console.Fatal("read from stdin failed")
			}
			content = string(bytes)
		}

		lines := strings.Split(content, "\n")
		if len(lines) == 0 {
			return
		}
		var (
			header string
			rows   []string
		)
		if noHeader {
			rows = lines
		} else {
			header, rows = lines[0], lines[1:]
		}

		if header != "" {
			tableWriter.AppendHeader(getRow(strings.Split(header, ","), TABLE_COLORS[color]))
		}

		for _, line := range rows {
			if line == "" {
				continue
			}
			tableWriter.AppendRow(getRow(strings.Split(line, ",")))
		}
		fmt.Println(tableWriter.Render())
		tableWriter.Render()
	},
}

func init() {
	CSVRender.Flags().Bool("no-header", false, "Has table header")
	CSVRender.Flags().String("style", "default", "Table style. default, light, rounded, bold")
	CSVRender.Flags().String("color", "blue", "Table color")

}
