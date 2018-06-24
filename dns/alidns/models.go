package alidns

//Alidns represents Alidns attribute and method associates with it.
type Alidns struct {
}

// CreateDns to store all attribute to create Ali-cloud DNS
type CreateDns struct {
	DomainName string
	RR         string
	Type       string
	Value      string
	TTL        int
	Priority   int
	Line       string
}

// DeleteDns to store all attribute to delete Ali-cloud DNS
type DeleteDns struct {
	RecordId string
}

// ListResourceDnsRecordSets to store all attribute to list resource Ali-cloud DNS record sets
type ListDns struct {
	DomainName   string
	PageNumber   int
	PageSize     int
	RRKeyWord    string
	TypeKeyWord  string
	ValueKeyWord string
}
