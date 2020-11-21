package utils

import "os"

func GetHostName() string {
	name, _ := os.Hostname()
	return name
}
