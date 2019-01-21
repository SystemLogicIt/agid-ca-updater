package client

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
}

func NewClient(country string) *client {
	return &client{
		country: country,
	}
}
