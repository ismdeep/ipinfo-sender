package main

import "fmt"

func ShowHelpMsg() {
	fmt.Println("Usage: ipinfo-sender    -c <config.json path>")
	fmt.Println("    -c <config.json path>    config.json file path")
	fmt.Println()
	fmt.Println("    config.example.json shows at https://raw.githubusercontent.com/ismdeep/ipinfo-sender/main/config.example.json")
}
