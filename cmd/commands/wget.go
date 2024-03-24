package commands

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/BytemanD/easygo/pkg/global/logging"
	"github.com/BytemanD/easygo/pkg/progress"
	"github.com/BytemanD/easygo/pkg/stringutils"
	"github.com/spf13/cobra"
)

type ProgressWriter struct {
	Writer *bufio.Writer
	bar    *progress.ProgressBar
}

func (pw ProgressWriter) Write(p []byte) (n int, err error) {
	pw.bar.Increment(len(p))
	return pw.Writer.Write(p)
}
func (pw ProgressWriter) Flush() error {
	return pw.Writer.Flush()
}

var Wget = &cobra.Command{
	Use:   "wget",
	Short: "get web file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		output, _ := cmd.Flags().GetString("output")

		resp, err := http.Get(url)
		if err != nil {
			logging.Error("get %s failed: %s", url, err)
			return
		}
		size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
		_, fileName := filepath.Split(url)
		err = os.MkdirAll(output, os.ModePerm)
		if err != nil {
			logging.Error("make dir %s failed: %s", outputPath, err)
			return
		}
		outputPath := path.Join(output, fileName)
		outputFile, err := os.Create(outputPath)
		if err != nil {
			logging.Error("create %s failed: %s", outputPath, err)
			return
		}
		defer outputFile.Close()

		logging.Info("size: %s", stringutils.HumanBytes(size))
		logging.Info("saving to %s", outputPath)
		writer := progress.NewProgressWriter(outputFile, size)
		io.Copy(writer, resp.Body)
	},
}

func init() {
	Wget.Flags().StringP("output", "O", "./", "保存路径")

}
