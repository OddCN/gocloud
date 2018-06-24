package vultrdns

import (
	"github.com/cloudlibz/gocloud/vultrauth"
	"testing"
)

func init() {
	vultrauth.LoadConfig()
}

func TestVultrDNS_CreateDns(t *testing.T) {
	var vultrDNS VultrDNS
	CreateDns := map[string]interface{}{
		"domain": "oddcn.cn",
		"name":   "gocloud.test",
		"type":   "A",
		"data":   "192.0.2.1",
	}
	resp, err := vultrDNS.CreateDns(CreateDns)
	if err != nil {
		t.Errorf("CreateDns Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr DNS record is created successfully.")
}

func TestVultrDNS_ListDns(t *testing.T) {
	var vultrDNS VultrDNS
	ListDns := map[string]interface{}{
		"domain": "oddcn.cn",
	}
	resp, err := vultrDNS.ListDns(ListDns)
	if err != nil {
		t.Errorf("ListDns Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr DNS record is listed: %s", response["body"])
}

func TestVultrDNS_DeleteDns(t *testing.T) {
	var vultrDNS VultrDNS
	DeleteDns := map[string]interface{}{
		"domain":   "oddcn.cn",
		"RECORDID": 7065075,
	}
	resp, err := vultrDNS.DeleteDns(DeleteDns)
	if err != nil {
		t.Errorf("DeleteDns Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr DNS record is deleted")
}
