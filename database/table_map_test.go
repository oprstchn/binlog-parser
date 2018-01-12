// +build integration

package database

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)


func TestLookupTableMetadata(t *testing.T) {
	const SCHEMA = "employess"
	db, _ := GetDatabaseInstance(TEST_DB_DSN_STRING)
	defer db.Close()
	t.Log(db);

	t.Run("Found", func(t *testing.T) {
		tableMap := NewTableMap(db)
		tableMap.Add(1, SCHEMA, "dept_emp")
		tableMap.Add(2, SCHEMA, "employees")

		assertTableMetadata(t, &tableMap, 1, SCHEMA, "dept_emp")
		assertTableMetadata(t, &tableMap, 2, SCHEMA, "employees")
		t.Log(tableMap)
	})

	t.Run("Fields", func(t *testing.T) {
		tableMap := NewTableMap(db)
		tableMap.Add(1, SCHEMA, "buildings")

		tableMetadata, ok := tableMap.LookupTableMetadata(1)

		if ok != true {
			t.Fatal("Expected table metadata to be found")
		}

		expectedFields := map[int]string{
			0: "emp_no",
			1: "dept_no",
			2: "from_date",
			3: "to_date",
		}

		if !reflect.DeepEqual(tableMetadata.Fields, expectedFields) {
			t.Fatal("Wrong fields in table metadata")
		}
	})

	t.Run("Not Found", func(t *testing.T) {
		tableMap := NewTableMap(db)
		_, ok := tableMap.LookupTableMetadata(999)

		if ok != false {
			t.Fatal("Expected table metadata not to be found")
		}
	})

}

func assertTableMetadata(t *testing.T, tableMap *TableMap, tableId uint64, expectedSchema string, expectedTable string) {
	tableMetadata, ok := tableMap.LookupTableMetadata(tableId)

	if ok != true {
		t.Fatal(fmt.Sprintf("metadata not found for table id %d", tableId))
	}

	if tableMetadata.Schema != expectedSchema {
		t.Fatal(fmt.Sprintf("wrong schema name for table id %d", tableId))
	}

	if tableMetadata.Table != expectedTable {
		t.Fatal(fmt.Sprintf("wrong table name for table id %d", tableId))
	}
}
