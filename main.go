package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/google/gopacket/pcap"
)

func main() {

	path := flag.String("f", "ifdevs.csv", "file path")
	flag.Parse()

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

	f, err := os.Create(*path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	fmt.Println(" 💾 Saving interfaces...")

	if err := report.Execute(f, ds); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf(" 👍 Interfaces saved in file %q\n", *path)
}
