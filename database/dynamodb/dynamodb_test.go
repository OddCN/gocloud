package dynamodb

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"

func init() {
	awsAuth.LoadConfig()
}

func TestDescribeTables(t *testing.T) {

	var dynamodb Dynamodb

	DescribeTables := map[string]interface{}{
		"Region":    "us-east-2",
		"TableName": "hello",
	}

	_, err := bigtable.DescribeTables(DescribeTables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListTables(t *testing.T) {

	var dynamodb Dynamodb

	ListTables := map[string]interface{}{
		"Region":    "us-east-2",
		"TableName": "hello",
	}

	_, err := bigtable.ListTables(ListTables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeleteTables(t *testing.T) {

	var dynamodb Dynamodb

	DeleteTables := map[string]interface{}{
		"Region":    "us-east-2",
		"TableName": "hello",
	}

	_, err := dynamodb.DeleteTables(DeleteTables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestCreateTables(t *testing.T) {

	var dynamodb Dynamodb

	keySchema := []map[string]interface{}{
		{
			"AttributeName": "ForumName",
			"KeyType":       "HASH",
		},
		{
			"AttributeName": "Subject",
			"KeyType":       "RANGE",
		},
	}

	attributeDefinitions := []map[string]interface{}{

		{
			"AttributeName": "ForumName",
			"AttributeType": "S",
		},
		{
			"AttributeName": "Subject",
			"AttributeType": "S",
		},
		{
			"AttributeName": "LastPostDateTime",
			"AttributeType": "S",
		},
	}

	projection := map[string]interface{}{
		"ProjectionType": "KEYS_ONLY",
	}

	provisionedThroughput := map[string]interface{}{
		"ReadCapacityUnits":  5,
		"WriteCapacityUnits": 5,
	}

	localSecondaryIndexes := []map[string]interface{}{
		{
			"IndexName":  "LastPostIndex",
			"KeySchema":  keySchema,
			"Projection": projection,
		},
	}

	CreateTables := map[string]interface{}{
		"Region":                "us-east-1",
		"TableName":             "Thread",
		"KeySchema":             keySchema,
		"AttributeDefinitions":  attributeDefinitions,
		"LocalSecondaryIndexes": localSecondaryIndexes,
		"ProvisionedThroughput": provisionedThroughput,
	}

	_, err := dynamodb.CreateTables(CreateTables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
