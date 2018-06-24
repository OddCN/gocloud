package bigtable

import "testing"
import "fmt"

//create bigtable Createtable
func TestCreateTables(t *testing.T) {

	var bigtable Bigtable

	table := make(map[string]interface{})

	initialSplits := make([]map[string]interface{}, 0)

	CreateTables := map[string]interface{}{
		"parent":        "projects/adept-comfort-202709/instances/helloo",
		"tableId":       "tableId",
		"table":         table,
		"initialSplits": initialSplits,
	}

	_, err := bigtable.CreateTables(CreateTables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDescribeTables(t *testing.T) {

	var bigtable Bigtable

	DescribeTables := map[string]string{
		"name": "projects/adept-comfort-202709/instances/helloo/tables/bokkkya",
	}

	_, err := bigtable.DescribeTables(DescribeTables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListTables(t *testing.T) {

	var bigtable Bigtable

	ListTables := map[string]string{
		"parent":    "projects/adept-comfort-202709/instances/helloo",
		"view":      "NAME_ONLY",
		"pageToken": "",
	}

	_, err := bigtable.ListTables(ListTables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeleteTables(t *testing.T) {

	var bigtable Bigtable

	DeleteTables := map[string]string{
		"name": "projects/adept-comfort-202709/instances/helloo/tables/bokkkya",
	}

	_, err := bigtable.DeleteTables(DeleteTables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
