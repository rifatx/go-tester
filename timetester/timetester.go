package timetester

import (
	"fmt"
	"time"
)

const (
	TIME_FORMAT = "20060102T150405Z0700"
)

func Test() {
	t := time.Now()

	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(time.Now().Format(TIME_FORMAT))
}
