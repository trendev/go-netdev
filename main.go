package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/google/gopacket/pcap"
)

func main() {
	fmt.Println(" 🔍 Collecting devices details...")

	ds, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalln(err)
	}

	const templ = `id;name;IP;mask
{{range $i, $d := .}}{{$name := .Name}}{{range .Addresses}}{{$i}};{{$name}};{{.IP}};{{.Netmask}}
{{else}}{{$i}};{{$name}};;
{{end}}{{end}}`

	report := template.Must(template.New("templ").Parse(templ))

	path := "ifdevs.csv"

	f, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	fmt.Println(" 💾 Saving interfaces...")

	if err := report.Execute(f, ds); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf(" 👍 Interfaces saved in file %q\n", path)
}
