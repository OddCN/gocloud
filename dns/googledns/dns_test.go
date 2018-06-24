package googledns

import "testing"

func TestCreateDns(t *testing.T) {
	var googledns Googledns

	CreateDns := map[string]interface{}{
		"Project":     "sheltermap-1493101612061",
		"Kind":        "dns#managedZone",
		"Description": "dns",
		"DnsName":     "rootmonk.me.",
		"Name":        "gocloud",
	}
	_, err := googledns.CreateDns(CreateDns)

	if err != nil {
		t.Errorf("Test Fail")
	}

}

func TestListDns(t *testing.T) {
	var googledns Googledns

	ListDns := map[string]string{
		"Project": "sheltermap-1493101612061",
	}

	_, err := googledns.ListDns(ListDns)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListResourceDnsRecordSets(t *testing.T) {
	var googledns Googledns

	ListResourceDnsRecordSets := map[string]string{
		"Project":     "sheltermap-1493101612061",
		"managedZone": "gocloud3",
	}
	_, err := googledns.ListResourceDnsRecordSets(ListResourceDnsRecordSets)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeleteDns(t *testing.T) {
	var googledns Googledns

	DeleteDns := map[string]string{
		"Project":     "sheltermap-1493101612061",
		"managedZone": "gocloud3",
	}
	_, err := googledns.DeleteDns(DeleteDns)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
