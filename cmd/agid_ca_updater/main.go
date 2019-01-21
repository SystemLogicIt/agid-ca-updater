package main

import (
	"flag"
	"github.com/systemlogic/agid-ca-updater/client"
	"log"
)

var country = flag.String("country", "IT", "country code to fetch")

func main() {
	flag.Parse()

	client := client.NewClient(*country)

	log.Printf(`[+] Downloading remote file from: %s\n`, client)

}
