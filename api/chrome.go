// API for third-party servers, ChromeDB in our case
// Body example:
//"<?xml version=\"1.0\" encoding=\"utf-8\"?>" +
//	"<soapenv:Envelope xmlns:soapenv=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:urn=\"urn:description7b.services.chrome.com\">" +
//		"<soapenv:Header/>" +
//		"<soapenv:Body>" +
//			"<urn:TechnicalSpecificationDefinitionsRequest>" +
//				"<urn:accountInfo number=\"260420\" secret=\"0bb08aa63dee43eb\" country=\"US\" language=\"en\" />" +
//			"</urn:TechnicalSpecificationDefinitionsRequest>" +
//		"</soapenv:Body>" +
//	"</soapenv:Envelope>"
//
// You can do all jobs manually or use SOAP client like this https://github.com/justwatchcom/goat
package api

import (
	"decoder/models"
	"net/http"
	"bytes"
	"log"
	"io/ioutil"
	"fmt"
)

// Add here required calls
type ChromeApi interface {
	// Non-billable call, server health test
	CheckEndpoint() ()
	// Returns vin data by specified pattern
	Get(pattern string) (*models.Vin, error)
}

type ChromeHandler struct {
	Host string
	Port string
	Number int
	Secret string
	Country string
	Language string
	Request string
}

// Describe VinResponse here
type VinResponse struct {
	SomeField string
}

// Get VIN data
func (c *ChromeHandler) Get(pattern string) (*models.Vin, error) {
	// call to api endpoint
	rawResponse := c.makeSoapRequest("SomeGetVinDataEndpoint")
	vin, err := c.convertRawDataToVinModel(rawResponse)
	if (err != nil) {
		panic(fmt.Errorf("%v", "AAAAAAAAAAA PANIC!!!!!"))
	}

	return vin, nil
}

// Assemble request in soap format
func (c *ChromeHandler) composeRequest(endpoint string) *ChromeHandler {
	ch := c

	// compose here request body usign endpoint argument
	ch.Request = ""

	return ch
}

// Convert raw data to decoded VIN
func (c *ChromeHandler) convertRawDataToVinModel(raw VinResponse) (*models.Vin, error) {
	var m models.Vin
	return &m, nil
}

// Make request to endpoint and returns structure
func (c *ChromeHandler) makeSoapRequest(endpoint string) VinResponse {
	var response *http.Response
	var request *http.Request

	request, err := http.NewRequest("POST", c.Host, bytes.NewBuffer([]byte(c.Request)))
	if err == nil {
		response, err = (&http.Client{}).Do(request)
		defer response.Body.Close()
	} else {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	rsp := VinResponse{SomeField:string(b)}

	return rsp
}