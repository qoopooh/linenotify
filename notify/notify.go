package notify

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type SendOpts struct {
	Token   string // LINE notify token
	Message string // your message
	Prefix  string // prefix of a message with [squre brackets]
	Verbose bool   // print out debug message
}

func Send(opt SendOpts) string {

	const URL = "https://notify-api.line.me/api/notify"

	data := url.Values{}
	if opt.Prefix != "" {
		data.Set("message", fmt.Sprintf("[%s] %s", opt.Prefix, opt.Message))
	} else {
		data.Set("message", opt.Message)
	}

	if opt.Verbose {
		fmt.Printf("Send: %v\n", data)
	}

	req, _ := http.NewRequest(http.MethodPost,
		URL,
		strings.NewReader(data.Encode()))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", opt.Token))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(fmt.Sprintf("Err %s", err))
	}

	defer resp.Body.Close()
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	out := fmt.Sprintf("%v", res["message"])
	if out == "Missing Bearer" {
		out = fmt.Sprintf(
			"%s (Please set LINE_NOTIFY_TOKEN)", out)
	}

	if opt.Verbose {
		fmt.Printf("Resp: %v\n", res)
	}

	return out
}
