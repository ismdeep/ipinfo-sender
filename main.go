package main

import (
	"encoding/json"
	"fmt"
	"github.com/ismdeep/args"
	"github.com/ismdeep/ipinfo-sender/utils"
	"github.com/ismdeep/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Config struct {
	Host      string `json:"host"`
	Token     string `json:"token"`
	Client    string `json:"client"`
	Endless   bool   `json:"endless"`
	SleepTime int64  `json:"sleep_time"`
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func SendIpInfo(config *Config) {
	log.Debug("SendIpInfo()", "config", config)

	clientName := config.Client

	if clientName == "" {
		clientName = utils.GetHostName()
	}

	content := strings.Join(utils.GetIPAddressList(), "\n")

	params := url.Values{}
	params.Add("token", config.Token)
	params.Add("key", fmt.Sprintf("ipinfo-%s", clientName))
	params.Add("value", content)

	apiUrl := fmt.Sprintf("%s/api/status", config.Host)

	_, err := http.PostForm(apiUrl, params)
	if err != nil {
		log.Error("发送失败", "err", err)
	}
	log.Info("发送成功", "params", params)
}

func main() {
	if args.Exists("--version") {
		ShowVersionInfo()
		return
	}

	if args.Exists("--help") {
		ShowHelpMsg()
		return
	}

	if !args.Exists("-c") {
		ShowHelpMsg()
		return
	}

	config, err := LoadConfig(args.GetValue("-c"))
	if err != nil {
		fmt.Println(err)
		return
	}

	if config.Endless {
		for {
			SendIpInfo(config)
			time.Sleep(time.Duration(config.SleepTime) * time.Second)
		}
	} else {
		SendIpInfo(config)
	}
}
