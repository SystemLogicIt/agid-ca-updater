package client

import (
	"crypto/x509"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"io"
	"strings"
)

var countryURLs = map[string]string{
	`IT`: `https://eidas.agid.gov.it/TL/TSL-IT.xml`,
}

var serviceTypes = map[string]string{
	`CA/QC`:      `http://uri.etsi.org/TrstSvc/Svctype/CA/QC`,
	`IdV`:        `http://uri.etsi.org/TrstSvc/Svctype/IdV`,
	`TSA`:        `http://uri.etsi.org/TrstSvc/Svctype/TSA`,
	`TSA/QTST`:   `http://uri.etsi.org/TrstSvc/Svctype/TSA/QTST`,
	`TSA/TSS-QC`: `http://uri.etsi.org/TrstSvc/Svctype/TSA/TSS-QC`,
}

var serviceStatuses = map[string]string{
	`deprecated`: `http://uri.etsi.org/TrstSvc/TrustedList/Svcstatus/deprecatedatnationallevel`,
	`granted`:    `http://uri.etsi.org/TrstSvc/TrustedList/Svcstatus/granted`,
	`recognised`: `http://uri.etsi.org/TrstSvc/TrustedList/Svcstatus/recognisedatnationallevel`,
	`withdrawn`:  `http://uri.etsi.org/TrstSvc/TrustedList/Svcstatus/withdrawn`,
}

var serviceExtensions = map[string]string{
	`eSeals`:                `http://uri.etsi.org/TrstSvc/TrustedList/SvcInfoExt/ForeSeals`,
	`eSignatures`:           `http://uri.etsi.org/TrstSvc/TrustedList/SvcInfoExt/ForeSignatures`,
	`WebSiteAuthentication`: `http://uri.etsi.org/TrstSvc/TrustedList/SvcInfoExt/ForWebSiteAuthentication`,
}

type client struct {
	country string
	certs   TrustServiceStatusList
}

func NewClient(country string) *client {
	return &client{
		country: country,
		certs:   TrustServiceStatusList{},
	}
}

func (c client) URL() string {
	return countryURLs[c.country]
}

func (c *client) Decode(reader io.Reader) error {
	dec := xml.NewDecoder(reader)
	return dec.Decode(&c.certs)
}

func (c client) ExtractCertificates(x509Certs chan<- *x509.Certificate) error {
	defer close(x509Certs)

	for _, provider := range c.certs.TrustServiceProviderList.TrustServiceProvider {
		for _, srv := range provider.TSPServices.TSPService {
			xmlSvcInfo := srv.ServiceInformation

			if xmlSvcInfo.ServiceTypeIdentifier != serviceTypes[`CA/QC`] {
				continue
			}

			if xmlSvcInfo.ServiceStatus != serviceStatuses[`granted`] {
				continue
			}

			extensionFound := false
			for _, ext := range xmlSvcInfo.ServiceInformationExtensions.Extension {
				if ext.AdditionalServiceInformation.URI.Text == serviceExtensions[`eSignatures`] {
					extensionFound = true
					break
				}
			}

			if !extensionFound {
				continue
			}

			for _, digitalId := range xmlSvcInfo.ServiceDigitalIdentity.DigitalId {

				if digitalId.X509Certificate == `` {
					continue
				}

				pemText := "-----BEGIN CERTIFICATE-----\n"
				pemText += strings.Join(chunkString(digitalId.X509Certificate, 64), "\n")
				pemText += "\n-----END CERTIFICATE-----\n"

				block, rest := pem.Decode([]byte(pemText))
				if block == nil || len(rest) > 0 {
					return errors.New("Certificate decoding error")
				}
				cert, err := x509.ParseCertificate(block.Bytes)
				if err != nil {
					return err
				}

				x509Certs <- cert
			}
		}
	}

	return nil
}

func chunkString(s string, chunkSize int) []string {
	var chunks []string
	runes := []rune(s)

	if len(runes) == 0 {
		return []string{s}
	}

	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}
	return chunks
}
