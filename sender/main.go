package main

import (
	"fmt"
	"github.com/ismdeep/ipinfo-sender/sender/utils"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		ShowVersionInfo()
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		ShowHelpMsg()
		return
	}

	host := os.Getenv("MONITOR_HOST")
	token := os.Getenv("MONITOR_TOKEN")
	clientName := os.Getenv("MONITOR_CLIENT")
	if token == "" || host == "" {
		fmt.Println("Please set MONITOR_HOST and MONITOR_TOKEN in environment.")
		return
	}

	if clientName == "" {
		clientName = utils.GetHostName()
	}

	content := strings.Join(utils.GetIPAddressList(), "\n")

	params := url.Values{}
	params.Add("token", token)
	params.Add("key", fmt.Sprintf("ipinfo-%s", clientName))
	params.Add("value", content)

	apiUrl := fmt.Sprintf("%s/api/status", host)

	_, _ = http.PostForm(apiUrl, params)
}
