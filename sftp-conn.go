package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func (s *Servidor) run() (string, error) {

	fmt.Println("Conectando a ...." + s.Server)

	sshConfig := &ssh.ClientConfig{
		User: s.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(decrypt([]byte(keyEncrypt), s.Pass)),
		},
	}
	sshConfig.SetDefaults()

	sshConn := new(ssh.Client)
	var errx error

	for {
		sshConn, errx = ssh.Dial("tcp", s.Server, sshConfig)
		if errx == nil {
			break
		} else if strings.Index(errx.Error(), "unexpected message type 3") == -1 {
			return "Failed to dial: " + errx.Error(), errx
		} else {
			fmt.Print(".")
			time.Sleep(2 * time.Second)
		}
	}
	defer sshConn.Close()

	client, err := sftp.NewClient(sshConn)
	if err != nil {
		return "Failed to NewClient: " + err.Error(), err
	}
	defer client.Close()

	srcFile, err := client.Open(s.SrcFile)
	if err != nil {
		return "Failed to Open File: " + err.Error(), err
	}
	defer srcFile.Close()

	t := time.Now()
	ext := filepath.Ext(s.SrcFile)
	filenameDest := fmt.Sprintf("-%d-%02d-%02d-%02d-%02d-%02d%s", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), ext)
	filenameDest = fmt.Sprintf("%s%s%s", s.DstPath, s.SrcFile[:len(s.SrcFile)-len(ext)], filenameDest)

	dstFile, err := os.Create(filenameDest)
	if err != nil {
		return "Failed to Create local file: " + err.Error(), err
	}
	defer dstFile.Close()

	srcFile.WriteTo(dstFile)
	srcFile.Close()
	dstFile.Close()

	return fmt.Sprintf("OK: Servidor: %s  Destino: %s  ", s.Server, filenameDest), nil
}

func runServer(params []string) string {

	if len(params) != 3 {
		fmt.Println("Debe de enviar como parámetro  --all o bien el número de configuración, utilice list")
		return ""
	} else if indice, err := strconv.Atoi(params[2]); err == nil || strings.EqualFold(params[2], "--all") {

		v := Configuracion{}
		//Lee el archivo de configuración
		openXML(&v)

		srvlist := []Servidor{}

		if err == nil {
			srvlist = append(srvlist, v.Servidores[indice])
		} else {
			srvlist = v.Servidores
		}

		for indice, servidor := range srvlist {
			str, _ := servidor.run()
			fmt.Printf("[%d] %s \n", indice, str)
		}
	}

	return "Proceso terminado...."
}

func showVersion() string {
	return "Versión de prueba para PEPSICO"
}

func showExpiredTrial() string {
	return "Por favor comunicarse con hector@globalsiag.com para validar la versión de prueba"
}
