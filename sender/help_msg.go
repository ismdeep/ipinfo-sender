package main

import "fmt"

func ShowHelpMsg() {
	fmt.Println("Usage: ipinfo-sender    -c <config.json path>    [-t]")
	fmt.Println("    -t    Loop forever")
	fmt.Println("    -c <config.json path>    config.json file path")
}
