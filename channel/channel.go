package channel

import (
	"fmt"
	"strconv"
)

func Test() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- "Hello " + strconv.Itoa(i)
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
	//ch<-"lkj" //panic: closed channel
}
