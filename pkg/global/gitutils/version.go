package gitutils

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/fjboy/magic-pocket/pkg/global/logging"
)

func getOutput(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		logging.Error("命令 '%s' 执行失败: %s", cmd, err)
		return "", err
	}
	return string(out), err
}

func getCommitNum(startTag string, endTag string) int {
	tagRange := startTag
	if endTag != "" {
		tagRange += "..." + endTag
	}
	out, err := getOutput("git", "log", "--pretty=oneline", tagRange)
	if err != nil {
		return 0
	}
	commits := strings.Split(strings.TrimSpace(out), "\n")
	if commits[0] == "" {
		return len(commits[1:])
	} else {
		return len(commits)
	}
}

func GetVersion() string {
	re := regexp.MustCompile("^[vV]*[0-9]")
	out, err := getOutput("git", "tag", "--sort=-taggerdate")
	var (
		lastTag string
		nums    int
	)
	if err == nil {
		tags := strings.Split(out, "\n")
		for i := 0; i < len(tags); i++ {
			if re.FindString(tags[i]) != "" {
				lastTag = tags[i]
				break
			}
		}
	}
	nums = getCommitNum("HEAD", lastTag)
	if lastTag == "" {
		lastTag = "0.0.0"
	}
	if nums == 0 {
		return lastTag
	} else {
		return fmt.Sprintf("%s.dev%d", lastTag, nums)
	}
}
