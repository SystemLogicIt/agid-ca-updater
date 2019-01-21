package main

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"flag"
	"fmt"
	"github.com/SystemLogicIt/agid-ca-updater/client"
	"log"
	"net/http"
	"os"
)

// Extra ASN1 OIDs that we may need to handle
var (
	oidEmailAddress                 = []int{1, 2, 840, 113549, 1, 9, 1}
	oidExtensionAuthorityInfoAccess = []int{1, 3, 6, 1, 5, 5, 7, 1, 1}
	oidNSComment                    = []int{2, 16, 840, 1, 113730, 1, 13}
)

var country = flag.String("country", "IT", "country code to fetch")

func main() {
	flag.Parse()

	cli := client.NewClient(*country)

	log.Printf("[+] Downloading remote file from: %s\n", cli.URL())

	certsChan := make(chan *x509.Certificate, 100)

	go func() {
		// Get the data
		resp, err := http.Get(cli.URL())
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		err = cli.Decode(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = cli.ExtractCertificates(certsChan)
		if err != nil {
			log.Fatal(err)
		}
	}()

	outputFile, err := os.Create(`ca-bundle-it.crt`)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	for cert := range certsChan {

		log.Println(printName(cert.Subject.Names))

		block := &pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}
		err = pem.Encode(outputFile, block)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("[+] done\n")
}

func printName(names []pkix.AttributeTypeAndValue) []string {
	values := []string{}
	for _, name := range names {
		oid := name.Type
		if len(oid) == 4 && oid[0] == 2 && oid[1] == 5 && oid[2] == 4 {
			switch oid[3] {
			case 3:
				values = append(values, fmt.Sprintf("CN=%s", name.Value))
			case 6:
				values = append(values, fmt.Sprintf("C=%s", name.Value))
			case 8:
				values = append(values, fmt.Sprintf("ST=%s", name.Value))
			case 10:
				values = append(values, fmt.Sprintf("O=%s", name.Value))
			case 11:
				values = append(values, fmt.Sprintf("OU=%s", name.Value))
			default:
				values = append(values, fmt.Sprintf("UnknownOID=%s", name.Type.String()))
			}
		} else if oid.Equal(oidEmailAddress) {
			values = append(values, fmt.Sprintf("emailAddress=%s", name.Value))
		} else {
			values = append(values, fmt.Sprintf("UnknownOID=%s", name.Type.String()))
		}
	}

	return values
}
