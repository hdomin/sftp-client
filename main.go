package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) > 1 {
		if strings.EqualFold(os.Args[1], "add") {
			fmt.Println(addServer(os.Args))
		} else if strings.EqualFold(os.Args[1], "list") {
			fmt.Println(listServers())
		} else if strings.EqualFold(os.Args[1], "delete") {
			fmt.Println(deleteServer(os.Args))
		} else if strings.EqualFold(os.Args[1], "run") {
			fmt.Println(runServer(os.Args))
		}
	}
}
