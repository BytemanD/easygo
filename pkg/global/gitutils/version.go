package gitutils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/BytemanD/easygo/pkg/syscmd"
)

func getCommitNum(startTag string, endTag string) int {
	tagRange := startTag
	if endTag != "" {
		tagRange += "..." + endTag
	}
	out, err := syscmd.GetOutput("git", "log", "--pretty=oneline", tagRange)
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
	out, err := syscmd.GetOutput("git", "tag")
	var (
		lastTag string
		nums    int
	)
	if err == nil {
		tags := strings.Split(strings.Trim(out, "\n"), "\n")
		for i := range tags {
			tag := tags[len(tags)-1-i]
			if re.FindString(tag) != "" {
				lastTag = tag
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
