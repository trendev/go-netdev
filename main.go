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

	const templ = `id;name;IP;mask
{{range $i, $d := .}}{{$name := .Name}}{{range .Addresses}}{{$i}};{{$name}};{{.IP}};{{.Netmask}}
{{else}}{{$i}};{{$name}};;
{{end}}{{end}}`

	report := template.Must(template.New("templ").Parse(templ))

	if err := report.Execute(os.Stdout, ds); err != nil {
		log.Fatalln(err)
	}
}
