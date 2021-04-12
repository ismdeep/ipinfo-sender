package main

import (
	"encoding/json"
	"fmt"
	"github.com/ismdeep/ipinfo-sender/utils"
	"github.com/ismdeep/ismdeep-go-utils/args_util"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var logger *zap.Logger

func init() {
	logger = zap.NewExample()
}

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
	logger.Debug("SendIpInfo()", zap.Any("config", config))

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
		logger.Error("发送失败", zap.Any("err", err))
	}
	logger.Info("发送成功", zap.Any("params", params))
}

func main() {
	if args_util.Exists("--version") {
		ShowVersionInfo()
		return
	}

	if args_util.Exists("--help") {
		ShowHelpMsg()
		return
	}

	if !args_util.Exists("-c") {
		ShowHelpMsg()
		return
	}

	config, err := LoadConfig(args_util.GetValue("-c"))
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
