package twiliotester

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func SendSMS() {
	accountSid := "AC800b296a4f8a095146784e606005dfb1"
	authToken := "6a5676be4c203bd5d18856661e71fd75"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	v := url.Values{}
	//	v.Set("To", "+905355939369")
	//	v.Set("From", "14843712976")
	v.Set("Body", "Some test sms. pass: 123456")
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)

		a, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(string(a[:len(a)]))
	}
}
