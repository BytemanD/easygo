package syscmd

import (
	"os/exec"
	"strings"

	"github.com/BytemanD/easygo/pkg/global/logging"
)

func GetOutput(command string, args ...string) (string, error) {
	logging.Debug("Run: %s %s", command, strings.Join(args, " "))
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	logging.Debug("Output: %s, Error: %v", out, err)
	if err != nil {
		return "", err
	}
	return string(out), err
}
