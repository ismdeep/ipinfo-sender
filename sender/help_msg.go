package main

import "fmt"

func ShowHelpMsg() {
	fmt.Println("Usage: ")
	fmt.Println("    export MONITOR_HOST={HOST}")
	fmt.Println("    export MONITOR_TOKEN={TOKEN}")
	fmt.Println("    RUN ipinfo-sender")
}
