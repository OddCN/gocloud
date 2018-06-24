# gocloud DNS - gce

## Configure Google Cloud credentials.

Download your service account credentials file from Google Cloud and save it as `googlecloudinfo.json` in your <b>HOME/.gocloud</b> directory.

You can also set the credentials as environment variables:
```js
export PrivateKey =  "xxxxxxxxxxxx"
export Type =  "xxxxxxxxxxxx"
export ProjectID = "xxxxxxxxxxxx"
export PrivateKeyID = "xxxxxxxxxxxx"
export ClientEmail = "xxxxxxxxxxxx"
export ClientID = "xxxxxxxxxxxx"
export AuthURI = "xxxxxxxxxxxx"
export TokenURI = "xxxxxxxxxxxx"
export AuthProviderX509CertURL = "xxxxxxxxxxxx"
export ClientX509CertURL =  "xxxxxxxxxxxx"
```

## Initialize library

```js

import "github.com/cloudlibz/gocloud/gocloud"

googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)
```

### Create DNS

```js
CreateDns := map[string]interface{}{
		"Project":     "sheltermap-1493101612061",
		"Kind":        "dns#managedZone",
		"Description": "dns",
		"DnsName":     "rootmonk.me.",
		"Name":        "gocloud",
	}

  resp, err := googlecloud.CreateDns(CreateDns)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])

  ```

### List DNS

```js
  ListDns := map[string]string{
		"Project": "sheltermap-1493101612061",
	}

  resp, err := googlecloud.ListDns(ListDns)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Delete DNS

```js

  DeleteDns := map[string]string{
		"Project":     "sheltermap-1493101612061",
		"managedZone": "gocloud3",
	}

 resp, err := googlecloud.DeleteDns(DeleteDns)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Detach node with loadbalancer

```js
  ListResourceDnsRecordSets := map[string]string{
	"Project"     : "sheltermap-1493101612061",
	"managedZone" : "gocloud3",
  }

  resp, err := googlecloud.ListResourceDnsRecordSets(ListResourceDnsRecordSets)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```