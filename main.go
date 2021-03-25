package main

import (
	"html/template"
	"log"
	"os"

	"github.com/google/gopacket/pcap"
)

func main() {
	ds, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalln(err)
	}

	const templ = `Interfaces : {{ . | len }} 
#################
{{range $i, $d := .}}
[ {{$i}} ] - {{.Name}}{{range .Addresses}}
  IP :      {{.IP}}
  Netmask : {{.Netmask}}
  ------------------------------------------------------------{{end}}
{{end}}
`

	report := template.Must(template.New("templ").Parse(templ))

	if err := report.Execute(os.Stdout, ds); err != nil {
		log.Fatalln(err)
	}
}
