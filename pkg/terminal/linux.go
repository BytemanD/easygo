package terminal

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/BytemanD/easygo/pkg/syscmd"
	"github.com/BytemanD/go-console/console"
)

type Terminal struct {
	Columns int
	Lines   int
}

func getLinuxTerminal() *Terminal {
	out, err := syscmd.GetOutput("stty", "size")
	if err != nil {
		console.Warn("get terminal falied, %s %s", out, err)
		return nil
	}
	values := strings.Split(string(out), " ")
	if len(values) < 2 {
		return nil
	}
	lines, _ := strconv.Atoi(values[0])
	columns, _ := strconv.Atoi(values[1])
	return &Terminal{Columns: columns, Lines: lines}
}

func CurTerminal() *Terminal {
	switch runtime.GOOS {
	case "linux":
		return getLinuxTerminal()
	default:
		console.Warn("get terminal for %s is not supported", runtime.GOOS)
		return nil
	}
}
