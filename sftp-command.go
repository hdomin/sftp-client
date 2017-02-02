package main

import (
	"fmt"
	"os"
)

func main() {

	/*
		sshConfig := &ssh.ClientConfig{
			User: "xxx",
			Auth: []ssh.AuthMethod{
				ssh.Password("xxx"),
			},
		}

		sshConfig.SetDefaults()
		sshConn, err := ssh.Dial("tcp", "xxx", sshConfig)
		if err != nil {
			panic("Failed to dial: " + err.Error())
		}
		defer sshConn.Close()

		client, err := sftp.NewClient(sshConn)
		if err != nil {
			panic("Failed to NewClient: " + err.Error())
		}
		defer client.Close()

		srcFile, err := client.Open("permisos.xlsx")
		if err != nil {
			panic("Failed to Open File: " + err.Error())
		}
		defer srcFile.Close()

		dstFile, err := os.Create("permisos.xlsx")
		if err != nil {
			panic("Failed to Create local file: " + err.Error())
		}
		defer dstFile.Close()

		srcFile.WriteTo(dstFile)
		srcFile.Close()
		dstFile.Close()
	*/
	fmt.Println(os.Args)
	writeXML()
}
