package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const ConfigurationFileName string = "sftp-client.xml"

type Servidor struct {
	Server  string `xml:"server"`
	User    string `xml:"user"`
	Pass    string `xml:"pass"`
	SrcFile string `xml:"srcFile"`
	DstPath string `xml:"dstPath"`
}

type Configuracion struct {
	XMLName    xml.Name   `xml:"Configuracion"`
	Servidores []Servidor `xml:"Servidor"`
}

func (s *Configuracion) AddServidor(srv *Servidor) {
	s.Servidores = append(s.Servidores, *srv)
}

func (srv *Servidor) String() string {
	//return fmt.Sprintf("Servidor: %s   \t- Usuario: %s   \t- Origen: %s   \t- Destino: %s \t-Pass: %s", srv.Server, srv.User, srv.SrcFile, srv.DstPath, decrypt([]byte(keyEncrypt), srv.Pass))
	return fmt.Sprintf("Servidor: %s   \t- Usuario: %s   \t- Origen: %s   \t- Destino: %s ", srv.Server, srv.User, srv.SrcFile, srv.DstPath)
}

func (srv *Servidor) setParams(params []string) {
	length := len(params)

	for index, param := range params {
		param = strings.ToLower(param)

		if length > index+1 {
			switch param {
			case "--server":
				srv.Server = params[index+1]
			case "--user":
				srv.User = params[index+1]
			case "--pass":
				srv.Pass = encrypt([]byte(keyEncrypt), params[index+1])
			case "--srcfile":
				srv.SrcFile = params[index+1]
			case "--dstpath":
				srv.DstPath = params[index+1]
			}
		}
	}
}

func addServer(params []string) string {
	srv := Servidor{}
	srv.setParams(params)

	v := Configuracion{}
	//Lee el archivo de configuración
	openXML(&v)

	//Agrega el nuevo servidor
	v.AddServidor(&srv)

	writeXML(&v)

	return "...............added. " + srv.Server
}

func deleteServer(params []string) string {
	ret := "debe de ingresar el número de configuración, utilice  'list'"

	if len(params) > 2 {
		if indice, err := strconv.Atoi(params[2]); err == nil {

			v := Configuracion{}
			//Lee el archivo de configuración
			openXML(&v)

			if len(v.Servidores) > indice {
				v.Servidores = append(v.Servidores[:indice], v.Servidores[indice+1:]...)
				writeXML(&v)
				ret = fmt.Sprintf("..........Configuración [%d] eliminada", indice)
			} else {
				ret = "indice fuera de rango, utilice 'list'"
			}

		} else {
			ret = "debe de enviar un valor numérico"
		}

	}
	return ret
}

func openXML(v *Configuracion) {
	xmlFile, _ := os.Open(ConfigurationFileName)
	xmlData, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(xmlData, &v)
	/*
		fmt.Printf("Estructura: %#v\n", v)

		xmlString, err := xml.MarshalIndent(v, "", "    ")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%s \n", string(xmlString))
	*/
}

func writeXML(v *Configuracion) {
	/*
		xmlString, err := xml.MarshalIndent(v, "", "    ")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%s \n", string(xmlString))
	*/

	file, _ := os.Create(ConfigurationFileName)
	xmlWriter := io.Writer(file)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")

	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	file.Close()

}

func listServers() string {
	v := Configuracion{}
	//Lee el archivo de configuración
	openXML(&v)
	lista := ""

	for index, servidor := range v.Servidores {
		lista = lista + fmt.Sprintf("[%d] %s \n", index, servidor.String())
	}

	return string(lista)
}
