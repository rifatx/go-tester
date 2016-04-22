package concurrent

import (
	"fmt"
	"runtime"
	"time"
)

func Test() {
	dur1, _ := time.ParseDuration("10ms")

	runtime.GOMAXPROCS(2)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("String1 ", i)
			time.Sleep(dur1)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("String2 ", i)
			time.Sleep(dur1)
		}
	}()

	dur2, _ := time.ParseDuration("1s")
	time.Sleep(dur2)
}
