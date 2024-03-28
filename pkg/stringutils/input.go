package stringutils

import (
	"fmt"
	"strings"
)

func ScanfComfirm(message string, yes, no []string) bool {
	var confirm string
	for {
		fmt.Printf("%s [%s|%s]: ", message, strings.Join(yes, "/"), strings.Join(no, "/"))
		fmt.Scanf("%s %d %f", &confirm)
		if ContainsString(yes, confirm) {
			return true
		} else if ContainsString(no, confirm) {
			return false
		} else {
			fmt.Print("输入错误, 请重新输入!")
		}
	}
}
