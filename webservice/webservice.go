package webservice

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseLogger struct {
	http.ResponseWriter
}

func (w ResponseLogger) Write(b []byte) (int, error) {
	fmt.Println(string(b))
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

	log := func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
			handler.ServeHTTP(ResponseLogger{w}, r)
		})
	}

	http.ListenAndServe(":8080", log(http.DefaultServeMux))
}
