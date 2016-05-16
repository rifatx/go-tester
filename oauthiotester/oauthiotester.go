package oauthiotester

import (
	"fmt"
	"github.com/oauth-io/sdk-go"
	"net/http"
)

var PUBLIC_KEY = "1rAx-yg-DxaueF_YccNJIiS7B7Y"
var SECRET_KEY = "ESjihCvQy653OwfFL1233YRGDLY"

func Test() {
	//m := mux.NewRouter()
	oauth := oauthio.New(PUBLIC_KEY, SECRET_KEY)

	http.HandleFunc("/signin", oauth.Redirect("facebook", "http://localhost:8080/oauth/redirect"))

	http.HandleFunc("/oauth/redirect", oauth.Callback(func(res *oauthio.OAuthRequestObject, err error, rw http.ResponseWriter, req *http.Request) {
		if err != nil {
			fmt.Println(err)
			return
		}

		r, _ := res.Me([]string{})

		//fmt.Println(res.AccessToken, err)
		fmt.Println(string(r))
	}))

	http.ListenAndServe(":8080", http.DefaultServeMux)
}
