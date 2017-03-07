package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	end, _ := time.Parse(time.RFC822, "25 Nov 17 09:00 UTC")
	if time.Now().After(end) {
		fmt.Println(showExpiredTrial())
		return
	}

	if len(os.Args) > 1 {
		if strings.EqualFold(os.Args[1], "add") {
			fmt.Println(addServer(os.Args))
		} else if strings.EqualFold(os.Args[1], "list") {
			fmt.Println(listServers())
		} else if strings.EqualFold(os.Args[1], "delete") {
			fmt.Println(deleteServer(os.Args))
		} else if strings.EqualFold(os.Args[1], "run") {
			fmt.Println(runServer(os.Args))
		} else if strings.EqualFold(os.Args[1], "--version") {
			fmt.Println(showVersion())
		}
	}
}
