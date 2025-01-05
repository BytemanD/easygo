package syscmd

import (
	"os"
	"os/exec"
	"strings"

	"github.com/BytemanD/go-console/console"
)

func GetOutput(command string, args ...string) (string, error) {
	console.Debug("Run: %s %s", command, strings.Join(args, " "))
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	out, err := cmd.CombinedOutput()
	console.Debug("Output: %s, Error: %v", out, err)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), err
}
