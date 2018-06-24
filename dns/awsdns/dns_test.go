package awsdns

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"

func init() {
	awsAuth.LoadConfig()
}

func TestCreateDns(t *testing.T) {
	var awsdns Awsdns
	CreateDns := map[string]interface{}{
		"name":             "rootmonk.me",
		"hostedZoneConfig": "hostedZoneConfig",
	}

	awsdns.CreateDns(CreateDns)
}

func TestDeleteDns(t *testing.T) {
	var awsdns Awsdns
	DeleteDns := map[string]string{
		"ID": "ZOD7SUP0ZGGQQ",
	}

	_, err := awsdns.DeleteDns(DeleteDns)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListDns(t *testing.T) {
	var awsdns Awsdns

	ListDns := map[string]interface{}{
		"marker":   "",
		"maxItems": 2,
	}

	_, err := awsdns.ListDns(ListDns)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListResourceDnsRecordSets(t *testing.T) {
	var awsdns Awsdns
	ListResourceDnsRecordSets := map[string]interface{}{
		"zone": "ZBNX5TIW033J2",
	}

	_, err := awsdns.ListResourceDnsRecordSets(ListResourceDnsRecordSets)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
