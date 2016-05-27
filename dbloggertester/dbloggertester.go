package dbloggertester

import (
	"database/sql"
	"io"
	"log"
	"os"
	"time"
)

type DbWriter struct {
	driverName      string
	dataSourceName  string
	failsafeLogFile string
}

func Init(driverName string, dataSourceName string, failsafeLogFile string) *log.Logger {
	var w io.Writer = DbWriter{driverName: driverName, dataSourceName: dataSourceName, failsafeLogFile: failsafeLogFile}

	return log.New(w, "[DB] ", 0)
}

func (dbw DbWriter) Write(p []byte) (int, error) {
	go innerWrite(&dbw, p)
	return 1, nil
}

func innerWrite(dbw *DbWriter, p []byte) {
	db, err := sql.Open(dbw.driverName, dbw.dataSourceName)
	if err == nil {
		defer db.Close()

		_, err = db.Exec("INSERT INTO LOG.LOGS(date, message) VALUES (?, ?)", time.Now, p)
	}

	if err != nil {
		f, err := os.OpenFile(dbw.failsafeLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			f = os.Stderr
		}

		logger := log.New(f, "[FILE] ", log.Ldate|log.Ltime)
		logger.Fatalf("could not log to db, writing to file %s: %s", f.Name(), p)
	}
}

func Test() {
	dblogger := Init("sd", "sdf", "D:\\Desktop\\golog.log")

	dblogger.Println("osman")

	//simulate a running service
	time.Sleep(1000)
}
