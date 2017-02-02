package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Servidor struct {
	Server  string `xml:"server"`
	User    string `xml:"user"`
	Pass    string `xml:"pass"`
	SrcFile string `xml:"srcFile"`
	DstPath string `xml:"dstPath"`
}

type Configuracion struct {
	XMLName    xml.Name `xml:"configuracion"`
	Servidores ServidorArray
}

type ServidorArray struct {
	Servidores []Servidor
}

func (s *ServidorArray) AddServidor(pserver string, puser string, ppass string, psrcFile string, pdstPath string) {
	fmt.Println(pserver)
	srv := Servidor{Server: pserver, User: puser, Pass: ppass, SrcFile: psrcFile, DstPath: pdstPath}
	s.Servidores = append(s.Servidores, srv)
}

func writeXML() {
	v := Configuracion{}
	v.Servidores.AddServidor("servidor1:22", "user1", "pass1", "file-origen1", "file-destino1")
	v.Servidores.AddServidor("servidor2:22", "user2", "pass2", "file-origen2", "file-destino2")
	v.Servidores.AddServidor("servidor3:22", "user3", "pass3", "file-origen3", "file-destino3")

	xmlString, err := xml.MarshalIndent(v, "", "    ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s \n", string(xmlString))

	// everything ok now, write to file.
	filename := "newstaffs.xml"
	file, _ := os.Create(filename)

	xmlWriter := io.Writer(file)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

}

// https://www.socketloop.com/tutorials/golang-create-new-xml-file
