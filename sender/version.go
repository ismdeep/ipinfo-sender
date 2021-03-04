package main

import (
	"fmt"
)

const VERSION = "0.0.2"

func GetVersion() string {
	return VERSION
}

func ShowVersionInfo() {
	fmt.Println(GetVersion())
}
