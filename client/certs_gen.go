// TrustServiceStatusList was generated 2019-01-21 21:06:42 by GENESIOMBP\genesio on GENESIOMBP.
package client

import "encoding/xml"

type TrustServiceStatusList struct {
	XMLName           xml.Name `xml:"TrustServiceStatusList"`
	Text              string   `xml:",chardata"`
	Xmlns             string   `xml:"xmlns,attr"`
	Ns2               string   `xml:"ns2,attr"`
	Ns3               string   `xml:"ns3,attr"`
	Ns4               string   `xml:"ns4,attr"`
	Ns5               string   `xml:"ns5,attr"`
	Ns6               string   `xml:"ns6,attr"`
	ID                string   `xml:"Id,attr"`
	TSLTag            string   `xml:"TSLTag,attr"`
	SchemeInformation struct {
		Text                 string `xml:",chardata"`
		TSLVersionIdentifier string `xml:"TSLVersionIdentifier"`
		TSLSequenceNumber    string `xml:"TSLSequenceNumber"`
		TSLType              string `xml:"TSLType"`
		SchemeOperatorName   struct {
			Text string `xml:",chardata"`
			Name []struct {
				Text string `xml:",chardata"`
				Lang string `xml:"lang,attr"`
			} `xml:"Name"`
		} `xml:"SchemeOperatorName"`
		SchemeOperatorAddress struct {
			Text            string `xml:",chardata"`
			PostalAddresses struct {
				Text          string `xml:",chardata"`
				PostalAddress []struct {
					Text          string `xml:",chardata"`
					Lang          string `xml:"lang,attr"`
					StreetAddress string `xml:"StreetAddress"`
					Locality      string `xml:"Locality"`
					PostalCode    string `xml:"PostalCode"`
					CountryName   string `xml:"CountryName"`
				} `xml:"PostalAddress"`
			} `xml:"PostalAddresses"`
			ElectronicAddress struct {
				Text string `xml:",chardata"`
				URI  []struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
				} `xml:"URI"`
			} `xml:"ElectronicAddress"`
		} `xml:"SchemeOperatorAddress"`
		SchemeName struct {
			Text string `xml:",chardata"`
			Name []struct {
				Text string `xml:",chardata"`
				Lang string `xml:"lang,attr"`
			} `xml:"Name"`
		} `xml:"SchemeName"`
		SchemeInformationURI struct {
			Text string `xml:",chardata"`
			URI  []struct {
				Text string `xml:",chardata"`
				Lang string `xml:"lang,attr"`
			} `xml:"URI"`
		} `xml:"SchemeInformationURI"`
		StatusDeterminationApproach string `xml:"StatusDeterminationApproach"`
		SchemeTypeCommunityRules    struct {
			Text string `xml:",chardata"`
			URI  []struct {
				Text string `xml:",chardata"`
				Lang string `xml:"lang,attr"`
			} `xml:"URI"`
		} `xml:"SchemeTypeCommunityRules"`
		SchemeTerritory     string `xml:"SchemeTerritory"`
		PolicyOrLegalNotice struct {
			Text           string `xml:",chardata"`
			TSLLegalNotice []struct {
				Text string `xml:",chardata"`
				Lang string `xml:"lang,attr"`
			} `xml:"TSLLegalNotice"`
		} `xml:"PolicyOrLegalNotice"`
		HistoricalInformationPeriod string `xml:"HistoricalInformationPeriod"`
		PointersToOtherTSL          struct {
			Text            string `xml:",chardata"`
			OtherTSLPointer struct {
				Text                     string `xml:",chardata"`
				ServiceDigitalIdentities struct {
					Text                   string `xml:",chardata"`
					ServiceDigitalIdentity []struct {
						Text      string `xml:",chardata"`
						DigitalId struct {
							Text            string `xml:",chardata"`
							X509Certificate string `xml:"X509Certificate"`
						} `xml:"DigitalId"`
					} `xml:"ServiceDigitalIdentity"`
				} `xml:"ServiceDigitalIdentities"`
				TSLLocation           string `xml:"TSLLocation"`
				AdditionalInformation struct {
					Text             string `xml:",chardata"`
					OtherInformation []struct {
						Text               string `xml:",chardata"`
						TSLType            string `xml:"TSLType"`
						SchemeTerritory    string `xml:"SchemeTerritory"`
						MimeType           string `xml:"MimeType"`
						SchemeOperatorName struct {
							Text string `xml:",chardata"`
							Name []struct {
								Text string `xml:",chardata"`
								Lang string `xml:"lang,attr"`
							} `xml:"Name"`
						} `xml:"SchemeOperatorName"`
						SchemeTypeCommunityRules struct {
							Text string `xml:",chardata"`
							URI  struct {
								Text string `xml:",chardata"`
								Lang string `xml:"lang,attr"`
							} `xml:"URI"`
						} `xml:"SchemeTypeCommunityRules"`
					} `xml:"OtherInformation"`
				} `xml:"AdditionalInformation"`
			} `xml:"OtherTSLPointer"`
		} `xml:"PointersToOtherTSL"`
		ListIssueDateTime string `xml:"ListIssueDateTime"`
		NextUpdate        struct {
			Text     string `xml:",chardata"`
			DateTime string `xml:"dateTime"`
		} `xml:"NextUpdate"`
		DistributionPoints struct {
			Text string `xml:",chardata"`
			URI  string `xml:"URI"`
		} `xml:"DistributionPoints"`
	} `xml:"SchemeInformation"`
	TrustServiceProviderList struct {
		Text                 string `xml:",chardata"`
		TrustServiceProvider []struct {
			Text           string `xml:",chardata"`
			TSPInformation struct {
				Text    string `xml:",chardata"`
				TSPName struct {
					Text string `xml:",chardata"`
					Name []struct {
						Text string `xml:",chardata"`
						Lang string `xml:"lang,attr"`
					} `xml:"Name"`
				} `xml:"TSPName"`
				TSPTradeName struct {
					Text string `xml:",chardata"`
					Name []struct {
						Text string `xml:",chardata"`
						Lang string `xml:"lang,attr"`
					} `xml:"Name"`
				} `xml:"TSPTradeName"`
				TSPAddress struct {
					Text            string `xml:",chardata"`
					PostalAddresses struct {
						Text          string `xml:",chardata"`
						PostalAddress []struct {
							Text            string `xml:",chardata"`
							Lang            string `xml:"lang,attr"`
							StreetAddress   string `xml:"StreetAddress"`
							Locality        string `xml:"Locality"`
							PostalCode      string `xml:"PostalCode"`
							CountryName     string `xml:"CountryName"`
							StateOrProvince string `xml:"StateOrProvince"`
						} `xml:"PostalAddress"`
					} `xml:"PostalAddresses"`
					ElectronicAddress struct {
						Text string `xml:",chardata"`
						URI  []struct {
							Text string `xml:",chardata"`
							Lang string `xml:"lang,attr"`
						} `xml:"URI"`
					} `xml:"ElectronicAddress"`
				} `xml:"TSPAddress"`
				TSPInformationURI struct {
					Text string `xml:",chardata"`
					URI  []struct {
						Text string `xml:",chardata"`
						Lang string `xml:"lang,attr"`
					} `xml:"URI"`
				} `xml:"TSPInformationURI"`
			} `xml:"TSPInformation"`
			TSPServices struct {
				Text       string `xml:",chardata"`
				TSPService []struct {
					Text               string `xml:",chardata"`
					ServiceInformation struct {
						Text                  string `xml:",chardata"`
						ServiceTypeIdentifier string `xml:"ServiceTypeIdentifier"`
						ServiceName           struct {
							Text string `xml:",chardata"`
							Name struct {
								Text string `xml:",chardata"`
								Lang string `xml:"lang,attr"`
							} `xml:"Name"`
						} `xml:"ServiceName"`
						ServiceDigitalIdentity struct {
							Text      string `xml:",chardata"`
							DigitalId []struct {
								Text            string `xml:",chardata"`
								X509Certificate string `xml:"X509Certificate"`
								X509SubjectName string `xml:"X509SubjectName"`
								X509SKI         string `xml:"X509SKI"`
							} `xml:"DigitalId"`
						} `xml:"ServiceDigitalIdentity"`
						ServiceStatus                string `xml:"ServiceStatus"`
						StatusStartingTime           string `xml:"StatusStartingTime"`
						ServiceInformationExtensions struct {
							Text      string `xml:",chardata"`
							Extension []struct {
								Text                         string `xml:",chardata"`
								Critical                     string `xml:"Critical,attr"`
								ExpiredCertsRevocationInfo   string `xml:"ExpiredCertsRevocationInfo"`
								AdditionalServiceInformation struct {
									Text string `xml:",chardata"`
									URI  struct {
										Text string `xml:",chardata"`
										Lang string `xml:"lang,attr"`
									} `xml:"URI"`
								} `xml:"AdditionalServiceInformation"`
								Qualifications struct {
									Text                 string `xml:",chardata"`
									QualificationElement struct {
										Text       string `xml:",chardata"`
										Qualifiers struct {
											Text      string `xml:",chardata"`
											Qualifier struct {
												Text string `xml:",chardata"`
												URI  string `xml:"uri,attr"`
											} `xml:"Qualifier"`
										} `xml:"Qualifiers"`
										CriteriaList struct {
											Text      string `xml:",chardata"`
											Assert    string `xml:"assert,attr"`
											PolicySet []struct {
												Text             string `xml:",chardata"`
												PolicyIdentifier []struct {
													Text       string `xml:",chardata"`
													Identifier string `xml:"Identifier"`
												} `xml:"PolicyIdentifier"`
											} `xml:"PolicySet"`
											Description string `xml:"Description"`
											KeyUsage    struct {
												Text        string `xml:",chardata"`
												KeyUsageBit struct {
													Text string `xml:",chardata"`
													Name string `xml:"name,attr"`
												} `xml:"KeyUsageBit"`
											} `xml:"KeyUsage"`
										} `xml:"CriteriaList"`
									} `xml:"QualificationElement"`
								} `xml:"Qualifications"`
								TakenOverBy struct {
									Text string `xml:",chardata"`
									URI  struct {
										Text string `xml:",chardata"`
										Lang string `xml:"lang,attr"`
									} `xml:"URI"`
									TSPName struct {
										Text string `xml:",chardata"`
										Name []struct {
											Text string `xml:",chardata"`
											Lang string `xml:"lang,attr"`
										} `xml:"Name"`
									} `xml:"TSPName"`
									SchemeOperatorName struct {
										Text string `xml:",chardata"`
										Name []struct {
											Text string `xml:",chardata"`
											Lang string `xml:"lang,attr"`
										} `xml:"Name"`
									} `xml:"SchemeOperatorName"`
									SchemeTerritory string `xml:"SchemeTerritory"`
								} `xml:"TakenOverBy"`
							} `xml:"Extension"`
						} `xml:"ServiceInformationExtensions"`
					} `xml:"ServiceInformation"`
					ServiceHistory struct {
						Text                   string `xml:",chardata"`
						ServiceHistoryInstance []struct {
							Text                  string `xml:",chardata"`
							ServiceTypeIdentifier string `xml:"ServiceTypeIdentifier"`
							ServiceName           struct {
								Text string `xml:",chardata"`
								Name struct {
									Text string `xml:",chardata"`
									Lang string `xml:"lang,attr"`
								} `xml:"Name"`
							} `xml:"ServiceName"`
							ServiceDigitalIdentity struct {
								Text      string `xml:",chardata"`
								DigitalId []struct {
									Text            string `xml:",chardata"`
									X509SubjectName string `xml:"X509SubjectName"`
									X509SKI         string `xml:"X509SKI"`
								} `xml:"DigitalId"`
							} `xml:"ServiceDigitalIdentity"`
							ServiceStatus                string `xml:"ServiceStatus"`
							StatusStartingTime           string `xml:"StatusStartingTime"`
							ServiceInformationExtensions struct {
								Text      string `xml:",chardata"`
								Extension []struct {
									Text                         string `xml:",chardata"`
									Critical                     string `xml:"Critical,attr"`
									ExpiredCertsRevocationInfo   string `xml:"ExpiredCertsRevocationInfo"`
									AdditionalServiceInformation struct {
										Text string `xml:",chardata"`
										URI  struct {
											Text string `xml:",chardata"`
											Lang string `xml:"lang,attr"`
										} `xml:"URI"`
									} `xml:"AdditionalServiceInformation"`
									Qualifications struct {
										Text                 string `xml:",chardata"`
										QualificationElement struct {
											Text       string `xml:",chardata"`
											Qualifiers struct {
												Text      string `xml:",chardata"`
												Qualifier struct {
													Text string `xml:",chardata"`
													URI  string `xml:"uri,attr"`
												} `xml:"Qualifier"`
											} `xml:"Qualifiers"`
											CriteriaList struct {
												Text     string `xml:",chardata"`
												Assert   string `xml:"assert,attr"`
												KeyUsage struct {
													Text        string `xml:",chardata"`
													KeyUsageBit struct {
														Text string `xml:",chardata"`
														Name string `xml:"name,attr"`
													} `xml:"KeyUsageBit"`
												} `xml:"KeyUsage"`
												PolicySet struct {
													Text             string `xml:",chardata"`
													PolicyIdentifier struct {
														Text       string `xml:",chardata"`
														Identifier string `xml:"Identifier"`
													} `xml:"PolicyIdentifier"`
												} `xml:"PolicySet"`
												Description string `xml:"Description"`
											} `xml:"CriteriaList"`
										} `xml:"QualificationElement"`
									} `xml:"Qualifications"`
									TakenOverBy struct {
										Text string `xml:",chardata"`
										URI  struct {
											Text string `xml:",chardata"`
											Lang string `xml:"lang,attr"`
										} `xml:"URI"`
										TSPName struct {
											Text string `xml:",chardata"`
											Name []struct {
												Text string `xml:",chardata"`
												Lang string `xml:"lang,attr"`
											} `xml:"Name"`
										} `xml:"TSPName"`
										SchemeOperatorName struct {
											Text string `xml:",chardata"`
											Name []struct {
												Text string `xml:",chardata"`
												Lang string `xml:"lang,attr"`
											} `xml:"Name"`
										} `xml:"SchemeOperatorName"`
										SchemeTerritory string `xml:"SchemeTerritory"`
									} `xml:"TakenOverBy"`
								} `xml:"Extension"`
							} `xml:"ServiceInformationExtensions"`
						} `xml:"ServiceHistoryInstance"`
					} `xml:"ServiceHistory"`
				} `xml:"TSPService"`
			} `xml:"TSPServices"`
		} `xml:"TrustServiceProvider"`
	} `xml:"TrustServiceProviderList"`
	Signature struct {
		Text       string `xml:",chardata"`
		Ds         string `xml:"ds,attr"`
		ID         string `xml:"Id,attr"`
		SignedInfo struct {
			Text                   string `xml:",chardata"`
			CanonicalizationMethod struct {
				Text      string `xml:",chardata"`
				Algorithm string `xml:"Algorithm,attr"`
			} `xml:"CanonicalizationMethod"`
			SignatureMethod struct {
				Text      string `xml:",chardata"`
				Algorithm string `xml:"Algorithm,attr"`
			} `xml:"SignatureMethod"`
			Reference []struct {
				Text       string `xml:",chardata"`
				ID         string `xml:"Id,attr"`
				Type       string `xml:"Type,attr"`
				URI        string `xml:"URI,attr"`
				Transforms struct {
					Text      string `xml:",chardata"`
					Transform []struct {
						Text      string `xml:",chardata"`
						Algorithm string `xml:"Algorithm,attr"`
					} `xml:"Transform"`
				} `xml:"Transforms"`
				DigestMethod struct {
					Text      string `xml:",chardata"`
					Algorithm string `xml:"Algorithm,attr"`
				} `xml:"DigestMethod"`
				DigestValue string `xml:"DigestValue"`
			} `xml:"Reference"`
		} `xml:"SignedInfo"`
		SignatureValue struct {
			Text string `xml:",chardata"`
			ID   string `xml:"Id,attr"`
		} `xml:"SignatureValue"`
		KeyInfo struct {
			Text     string `xml:",chardata"`
			X509Data struct {
				Text            string `xml:",chardata"`
				X509Certificate string `xml:"X509Certificate"`
			} `xml:"X509Data"`
		} `xml:"KeyInfo"`
		Object struct {
			Text                 string `xml:",chardata"`
			QualifyingProperties struct {
				Text             string `xml:",chardata"`
				Xades            string `xml:"xades,attr"`
				Target           string `xml:"Target,attr"`
				SignedProperties struct {
					Text                      string `xml:",chardata"`
					ID                        string `xml:"Id,attr"`
					SignedSignatureProperties struct {
						Text               string `xml:",chardata"`
						SigningTime        string `xml:"SigningTime"`
						SigningCertificate struct {
							Text string `xml:",chardata"`
							Cert struct {
								Text       string `xml:",chardata"`
								CertDigest struct {
									Text         string `xml:",chardata"`
									DigestMethod struct {
										Text      string `xml:",chardata"`
										Algorithm string `xml:"Algorithm,attr"`
									} `xml:"DigestMethod"`
									DigestValue string `xml:"DigestValue"`
								} `xml:"CertDigest"`
								IssuerSerial struct {
									Text             string `xml:",chardata"`
									X509IssuerName   string `xml:"X509IssuerName"`
									X509SerialNumber string `xml:"X509SerialNumber"`
								} `xml:"IssuerSerial"`
							} `xml:"Cert"`
						} `xml:"SigningCertificate"`
					} `xml:"SignedSignatureProperties"`
					SignedDataObjectProperties struct {
						Text             string `xml:",chardata"`
						DataObjectFormat struct {
							Text            string `xml:",chardata"`
							ObjectReference string `xml:"ObjectReference,attr"`
							MimeType        string `xml:"MimeType"`
						} `xml:"DataObjectFormat"`
					} `xml:"SignedDataObjectProperties"`
				} `xml:"SignedProperties"`
			} `xml:"QualifyingProperties"`
		} `xml:"Object"`
	} `xml:"Signature"`
}
