# gocloud DNS - Ali-cloud

## Configure Ali-cloud credentials

Create `alicloudconfig.json` as follows,
```
{
  "AliAccessKeyID": "xxxxxxxxxxxx",
  "AliAccessKeySecret": "xxxxxxxxxxxx"
}
```

also You can setup environment variables as

```
export AliAccessKeyID =  "xxxxxxxxxxxx"
export AliAccessKeySecret = "xxxxxxxxxxxx"
```

## Initialize library

```
import "github.com/cloudlibz/gocloud/gocloud"

alicloud, _ := gocloud.CloudProvider(gocloud.Aliprovider)
```

### Create DNS

```
  CreateDns := map[string]interface{}{
  		"DomainName": "oddcn.cn",
  		"RR":         "test.gocloud",
  		"Type":       "A",
  		"Value":      "202.106.0.20",
  		"TTL":        600,
  		"Line":       "default",
  }

  resp, err := alicloud.CreateDns(CreateDns)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Delete DNS

```js
  DeleteDns := map[string]interface{}{
  		"RecordId": "3888946862348288",
  }

  resp, err := alicloud.DeleteDns(DeleteDns)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### List DNS

```js
  ListDns := map[string]interface{}{
  		"PageNumber": 1,
  		"PageSize":   20,
  }

  resp, err := alicloud.ListDns(ListDns)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### List resource DNS record sets

```j
  ListResourceDnsRecordSets := map[string]interface{}{
  		"DomainName": "oddcn.cn",
  }

  resp, err := alicloud.ListResourceDnsRecordSets(ListResourceDnsRecordSets)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```