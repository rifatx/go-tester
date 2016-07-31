package main

import (
	"fmt"
	"net/http"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

func checkToken(w http.ResponseWriter, r *http.Request) bool {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("_______________________________"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	err := jwtMiddleware.CheckJWT(w, r)

	if err != nil {
		fmt.Println(err)
	}

	return err == nil
}

func Start() {
	http.HandleFunc(
		"/checktoken",
		func(w http.ResponseWriter, r *http.Request) {
			if !checkToken(w, r) {
				fmt.Println("token error.")
			}
		})

	http.HandleFunc(
		"/tokentest",
		func(w http.ResponseWriter, r *http.Request) {
			c := http.Client{}

			req, _ := http.NewRequest("GET", "http://localhost:8080/checktoken", nil)
			req.Header = r.Header
			res, err := c.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			if res.StatusCode != http.StatusOK {
				fmt.Println(res.StatusCode)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})

	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func main() {
	Start()
}
