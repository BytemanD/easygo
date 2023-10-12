package syscmd

import (
	"os"
	"os/exec"
	"strings"

	"github.com/BytemanD/easygo/pkg/global/logging"
)

func GetOutput(command string, args ...string) (string, error) {
	logging.Debug("Run: %s %s", command, strings.Join(args, " "))
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	out, err := cmd.CombinedOutput()
	logging.Debug("Output: %s, Error: %v", out, err)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), err
}
