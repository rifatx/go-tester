package webservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	TIME_FORMAT = "20060102T150405Z0700"
)

type ResponseLogger struct {
	http.ResponseWriter
	f func(b []byte)
}

func (w ResponseLogger) Write(b []byte) (int, error) {
	w.f(b)
	return w.ResponseWriter.Write(b)
}

func Start() {
	type datatype struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Param string `json:"param"`
	}

	http.HandleFunc(
		"/testfunc.get",
		func(w http.ResponseWriter, r *http.Request) {
			d := datatype{Id: 1, Name: "osman", Param: r.URL.Query().Get("hede")}
			j, _ := json.Marshal(d)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(j)
		})

	http.HandleFunc(
		"/testfunc.post",
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			d := datatype{Id: 2, Name: "mahmut", Param: r.PostForm.Get("hede")}
			j, _ := json.Marshal(d)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(j)
		},
	)

	innerLogFunc := func(b []byte) {
		fmt.Fprintln("[%s] %s", time.Now().Format(TIME_FORMAT), string(b))
	}

	logFunc := func(b []byte) {
		go innerLogFunc(b)
	}

	log := func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logFunc([]byte(fmt.Sprintf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)))
			handler.ServeHTTP(ResponseLogger{ResponseWriter: w, f: logFunc}, r)
		})
	}

	http.ListenAndServe(":8080", log(http.DefaultServeMux))
}
