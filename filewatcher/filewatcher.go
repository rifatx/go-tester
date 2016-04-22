package filewatcher

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func Test() {
	type Invoice struct {
		Number string
		Amount float64
		PON    int
		Date   time.Time
	}

	const watchedPath = "D:\\Desktop\\files"

	mp := runtime.GOMAXPROCS(4)
	defer runtime.GOMAXPROCS(mp)

	for {
		d, _ := os.Open(watchedPath)
		files, _ := d.Readdir(-1)

		for _, fi := range files {
			filePath := watchedPath + "\\" + fi.Name()
			fmt.Printf("got file: %s\n", fi.Name())

			f, _ := os.Open(filePath)
			data, _ := ioutil.ReadAll(f)
			f.Close()
			os.Remove(filePath)

			go func(data string) {
				reader := csv.NewReader(strings.NewReader(data))
				records, _ := reader.ReadAll()
				for _, r := range records {
					invoice := new(Invoice)
					invoice.Number = r[0]
					invoice.Amount, _ = strconv.ParseFloat(r[1], 64)
					invoice.PON, _ = strconv.Atoi(r[2])
					unixTime, _ := strconv.ParseInt(r[3], 10, 64)
					invoice.Date = time.Unix(unixTime, 0)

					fmt.Printf("received inv: %v\n", invoice.Number)
				}
			}(string(data))
		}

		time.Sleep(50 * time.Millisecond)
	}
}
