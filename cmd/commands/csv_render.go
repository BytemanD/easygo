package commands

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/BytemanD/easygo/pkg/fileutils"
	"github.com/BytemanD/easygo/pkg/global/logging"
	"github.com/jedib0t/go-pretty/v6/table"

	// "github.com/BytemanD/easygo/pkg/table"
	"github.com/spf13/cobra"
)

func getRow(columns []string) table.Row {
	tableRow := table.Row{}
	for _, column := range columns {
		tableRow = append(tableRow, column)
	}
	return tableRow
}

var CSVRender = &cobra.Command{
	Use:   "csv-render [file path]",
	Short: "CSV格式渲染表格",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hasHeader, _ := cmd.Flags().GetBool("header")
		style, _ := cmd.Flags().GetString("style")

		tableWriter := table.NewWriter()
		switch style {
		case "rounded":
			tableWriter.SetStyle(table.StyleRounded)
		case "light":
			tableWriter.SetStyle(table.StyleLight)
		case "bold":
			tableWriter.SetStyle(table.StyleBold)
		case "default":
			tableWriter.SetStyle(table.StyleDefault)
		default:
			fmt.Println("invalid style:", style)
			os.Exit(1)
		}
		var (
			content string
			err     error
		)
		if len(args) != 0 {
			fp := fileutils.FilePath{Path: args[0]}
			content, err = fp.ReadAll()
			if err != nil {
				logging.Fatal("read from file %s failed", fp.Path)
			}
		} else {
			bytes, err := io.ReadAll(os.Stdin)
			if err != nil {
				logging.Fatal("read from stdin failed")
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
		if hasHeader {
			header, rows = lines[0], lines[1:]
		} else {
			rows = lines
		}

		if header != "" {
			tableWriter.AppendHeader(getRow(strings.Split(header, ",")))
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
	CSVRender.Flags().Bool("header", false, "Has table header")
	CSVRender.Flags().String("style", "default", "Table style. default, light, rounded, bold")

}
