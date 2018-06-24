```
package alidns
    import "github.com/cloudlibz/gocloud/dns/alidns"
```

### TYPES

```
type Alidns struct {
}
```
Alidns represents Alidns attribute and method associates with it.

```
func (alidns *Alidns) CreateDns(request interface{}) (resp interface{}, err error)
```
CreateDns add DNS record accept map[string]interface{}

```
func (alidns *Alidns) DeleteDns(request interface{}) (resp interface{}, err error)
```
DeleteDns delete DNS record accept map[string]interface{}

```
func (alidns *Alidns) ListResourceDnsRecordSets(request interface{}) (resp interface{}, err error)
```
ListResourceDnsRecordSets list resource DNS record sets accept map[string]interface{}

```
func (alidns *Alidns) ListDns(request interface{}) (resp interface{}, err error)
```
ListDns list DNS record accept map[string]interface{}

```
type CreateDns struct {
    DomainName string
    RR         string
    Type       string
    Value      string
    TTL        int
    Priority   int
    Line       string
}
```
CreateDns to store all attribute to create Ali-cloud DNS

```
type DeleteDns struct {
    RecordId string
}
```
DeleteDns to store all attribute to delete Ali-cloud DNS

```
type ListDns struct {
    PageNumber int
    PageSize   int
    KeyWord    string
    GroupId    string
}
```
ListDns to store all attribute to list Ali-cloud DNS

```
type ListResourceDnsRecordSets struct {
    DomainName   string
    PageNumber   int
    PageSize     int
    RRKeyWord    string
    TypeKeyWord  string
    ValueKeyWord string
}
```
ListResourceDnsRecordSets to store all attribute to list resource Ali-cloud DNS record sets