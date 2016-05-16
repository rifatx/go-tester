package configtester

import (
	"fmt"
	"github.com/minhajuddin/config"
)

var C struct {
	Host    string
	ENV     string
	DB      string
	Cache   string
	Websrvr struct {
		ApiURL string `yaml:"api_url"`
		Creds  struct {
			Username string
			Password string
		}
	}
}

func Test() {
	config.LoadFromFile("./configtester/config.yaml", &C, func(p ...interface{}) {})

	fmt.Println(C.Websrvr.ApiURL)
	fmt.Println(C.ENV)
	fmt.Println(C.Websrvr.Creds.Username)
}
