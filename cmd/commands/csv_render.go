package commands

import (
	"fmt"
	"strings"

	"github.com/BytemanD/easygo/pkg/fileutils"
	"github.com/BytemanD/easygo/stream"
	"github.com/BytemanD/go-console/console"
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/duke-git/lancet/v2/slice"
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
	Short: "将CSV格式文本渲染成表格(支持标准输入或者文件)",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
			return err
		}
		style, _ := cmd.Flags().GetString("style")
		color, _ := cmd.Flags().GetString("color")
		if style != "" {
			if _, ok := TABLE_STYLES[style]; !ok {
				return fmt.Errorf("invalid style: %s", style)
			}
		}
		if color != "" {
			if _, ok := TABLE_COLORS[color]; !ok {
				return fmt.Errorf("invalid color: %s", color)
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		noHeader, _ := cmd.Flags().GetBool("no-header")
		style, _ := cmd.Flags().GetString("style")
		color, _ := cmd.Flags().GetString("color")

		var (
			content string
			err     error
		)
		if len(args) != 0 {
			content, err = fileutils.ReadAll(args[0])
			if err != nil {
				console.Fatal("read from file %s failed", args[0])
			}
		} else {
			content, err = stream.ReadStringFromStdin()
			if err != nil {
				console.Fatal("read from stdin failed: %s", args[0])
			}
		}

		lines := strings.Split(content, "\n")
		if len(lines) == 0 {
			return
		}
		tableWriter := table.NewWriter()
		tableWriter.SetStyle(TABLE_STYLES[style])
		tableWriter.Style().Color.Border = text.Colors{TABLE_COLORS[color]}
		tableWriter.Style().Color.Separator = text.Colors{TABLE_COLORS[color]}

		appendRow := func(lines []string) {
			slice.ForEach(lines, func(_ int, line string) {
				if line == "" {
					return
				}
				tableWriter.AppendRow(getRow(strings.Split(line, ",")))
			})
		}

		if !noHeader {
			tableWriter.AppendHeader(getRow(strings.Split(lines[0], ","), TABLE_COLORS[color]))
			appendRow(lines[1:])
		} else {
			appendRow(lines)
		}
		fmt.Println(tableWriter.Render())
	},
}

func init() {
	CSVRender.Flags().Bool("no-header", false, "Has table header")
	CSVRender.Flags().String("style", "default", "Table style.\n"+
		strings.Join(maputil.Keys(TABLE_STYLES), ", "))
	CSVRender.Flags().String("color", "blue", "Table color:\n"+
		strings.Join(maputil.Keys(TABLE_COLORS), ", "))

}
