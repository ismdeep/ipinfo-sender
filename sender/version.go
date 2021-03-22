package main

import (
	"fmt"
)

const VERSION = "0.0.3"

func GetVersion() string {
	return VERSION
}

func ShowVersionInfo() {
	fmt.Println(GetVersion())
}
