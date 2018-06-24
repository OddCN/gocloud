# gocloud DNS - AWS

## Configure AWS credentials.

Create `amazoncloudconfig.json` in your <b>HOME/.gocloud</b> directory as follows:
```js
{
  "AWSAccessKeyID": "xxxxxxxxxxxx",
  "AWSSecretKey": "xxxxxxxxxxxx"
}
```

You can also set the credentials as environment variables:
```js
export AWSAccessKeyID =  "xxxxxxxxxxxx"
export AWSSecretKey = "xxxxxxxxxxxx"
```

## Initialize library

```js

import "github.com/cloudlibz/gocloud/gocloud"

amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)
```

### Create DNS

```js
  CreateDns := map[string]interface{}{
		"name":             "rootmonk.me",
		"hostedZoneConfig": "hostedZoneConfig",
	}

 resp, err := awsdns.CreateDns(CreateDns)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Delete DNS
```js
 DeleteDns := map[string]string{
		"ID": "ZOD7SUP0ZGGQQ",
	}

 resp, err := awsdns.DeleteDns(DeleteDns)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### List DNS

```js
 ListDns := map[string]interface{}{
		"marker":   "",
		"maxItems": 2,
	}

  resp, err := awsdns.ListDns(ListDns)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### List ResourcednsRecordSets

```js

 ListResourceDnsRecordSets := map[string]interface{}{
	"zone": "ZBNX5TIW033J2",
  }

 resp, err := awsdns.ListResourceDnsRecordSets(ListResourceDnsRecordSets)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```