package digioceandns

import (
	"bytes"
	"encoding/json"
	"fmt"
	digioceanAuth "github.com/cloudlibz/gocloud/digioceanauth"
	"io/ioutil"
	"net/http"
)

// dnsBasePath is the endpoint URL for digitalocean API.
const dnsBasePath = "https://api.digitalocean.com/v2/domains"

// CreateDns function creates a new DNS record.
func (digioceandns *Digioceandns) CreateDns(request interface{}) (resp interface{}, err error) {

	var dnsInstance Digioceandns                                     // Initialize LoadBalancer struct
	var domainName string                                            // To store domain name
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken // Fetch the DigiOceanAccessToken

	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	for key, value := range param {

		switch key {

		case "DomainName":
			domainNameValue, _ := value.(string)
			domainName = domainNameValue

		case "Type":
			typeDNS, _ := value.(string)
			dnsInstance.Type = typeDNS

		case "Name":
			name, _ := value.(string)
			dnsInstance.Name = name

		case "Data":
			data, _ := value.(string)
			dnsInstance.Data = data

		case "Priority":
			priority, _ := value.(int)
			dnsInstance.Priority = priority

		case "Port":
			port, _ := value.(int)
			dnsInstance.Port = port

		case "TimeToLive":
			ttl, _ := value.(int)
			dnsInstance.TimeToLive = ttl

		case "Weight":
			weight, _ := value.(int)
			dnsInstance.Weight = weight

		case "Flags":
			flags, _ := value.(int)
			dnsInstance.Flags = flags

		case "Tag":
			tag, _ := value.(string)
			dnsInstance.Tag = tag

		} // Closes switch-case

	} // Closes for loop

	dnsInstanceJSON, _ := json.Marshal(dnsInstance)
	dnsInstanceJSONString := string(dnsInstanceJSON)
	var dnsInstanceJSONStringbyte = []byte(dnsInstanceJSONString)

	url := dnsBasePath + "/" + domainName + "/records"

	CreateDnsReq, err := http.NewRequest("POST", url, bytes.NewBuffer(dnsInstanceJSONStringbyte))
	if err != nil {
		fmt.Println(err)
	}
	CreateDnsReq.Header.Set("Content-Type", "application/json")
	CreateDnsReq.Header.Set("Authorization", "Bearer "+DigiOceanAccessToken)

	CreateDnsResp, err := http.DefaultClient.Do(CreateDnsReq)
	if err != nil {
		fmt.Println(err)
	}

	defer CreateDnsResp.Body.Close()

	responseBody, err := ioutil.ReadAll(CreateDnsResp.Body)
	CreateDnsResponse := make(map[string]interface{})
	CreateDnsResponse["status"] = CreateDnsResp.StatusCode
	CreateDnsResponse["body"] = string(responseBody)
	resp = CreateDnsResponse

	return resp, err
}

// DeleteDns function deletes a DNS record.
func (digioceandns *Digioceandns) DeleteDns(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := dnsBasePath + "/" + options["DomainName"] + "/records/" + options["RecordID"]
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken // Fetch the DigiOceanAccessToken

	DeleteDnsReq, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	DeleteDnsReq.Header.Set("Content-Type", "application/json")
	DeleteDnsReq.Header.Set("Authorization", "Bearer "+DigiOceanAccessToken)

	DeleteDnsResp, err := http.DefaultClient.Do(DeleteDnsReq)
	if err != nil {
		fmt.Println(err)
	}

	defer DeleteDnsResp.Body.Close()

	responseBody, err := ioutil.ReadAll(DeleteDnsResp.Body)
	DeleteDnsResponse := make(map[string]interface{})
	DeleteDnsResponse["status"] = DeleteDnsResp.StatusCode
	DeleteDnsResponse["body"] = string(responseBody)
	resp = DeleteDnsResponse

	return resp, err
}

// ListDns function lists DNS records.
func (digioceandns *Digioceandns) ListDns(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := dnsBasePath + "/" + options["DomainName"] + "/records"
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken // Fetch the DigiOceanAccessToken

	ListDnsReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	ListDnsReq.Header.Set("Content-Type", "application/json")
	ListDnsReq.Header.Set("Authorization", "Bearer "+DigiOceanAccessToken)

	ListDnsResp, err := http.DefaultClient.Do(ListDnsReq)
	if err != nil {
		fmt.Println(err)
	}

	defer ListDnsResp.Body.Close()

	responseBody, err := ioutil.ReadAll(ListDnsResp.Body)
	ListDnsResponse := make(map[string]interface{})
	ListDnsResponse["status"] = ListDnsResp.StatusCode
	ListDnsResponse["body"] = string(responseBody)
	resp = ListDnsResponse

	return resp, err
}

// ListResourceDnsRecordSets function lists DNS record sets. DigitalOcean API
// doesn't provide functionality to suppport this function.
func (digioceandns *Digioceandns) ListResourceDnsRecordSets(request interface{}) (resp interface{}, err error) {

	return resp, err
}
