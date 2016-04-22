package regextester

import (
	"fmt"
	"regexp"
)

func Test() {
	re, _ := regexp.Compile(".*")

	fmt.Println(re.MatchString("osman"))
}
