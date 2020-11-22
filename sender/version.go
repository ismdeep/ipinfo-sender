package main

import (
	"fmt"
)

const VERSION = "0.0.1"

func GetVersion() string {
	return VERSION
}

func ShowVersionInfo() {
	fmt.Println(GetVersion())
}
