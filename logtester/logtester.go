package logtester

import (
	"log"
	"os"
)

func Test() {
	var l log.Logger
	f, err := os.OpenFile("D:\\Desktop\\golog.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		f = os.Stdout
	}

	l.SetFlags(log.Ldate | log.Ltime)
	l.SetPrefix("[ERROR] ")
	l.SetOutput(f)

	l.Fatalln("osman")
}
