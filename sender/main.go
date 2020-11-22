package main

import (
	"fmt"
	"github.com/ismdeep/ipinfo-sender/sender/utils"
	"net/http"
	"net/url"
	"os"
)

func main() {
	host := os.Getenv("MONITOR_HOST")
	token := os.Getenv("MONITOR_TOKEN")
	if token == "" || host == "" {
		fmt.Println("Please set MONITOR_TOKEN in environment.")
		return
	}

	//log.Printf("MONITOR_HOST: %s\n", host)
	//log.Printf("MONITOR_TOKEN: %s\n", token)

	content := ""
	ipList := utils.GetIPAddressList()
	for _, item := range ipList {
		content += item + "\n"
	}

	params := url.Values{}
	params.Add("token", token)
	params.Add("key", fmt.Sprintf("ipinfo-%s", utils.GetHostName()))
	params.Add("value", content)

	//log.Println(params.Encode())

	apiUrl := fmt.Sprintf("%s/api/status", host)

	_, _ = http.PostForm(apiUrl, params)

	//if err != nil {
	//	log.Fatalln(err.Error())
	//}

	//log.Println(resp.StatusCode)
	//log.Println(resp.Status)

	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//
	//fmt.Println(string(body))
}
