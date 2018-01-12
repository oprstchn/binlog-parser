package database

import (
	"fmt"
	"reflect"
	"testing"
)

const (
	Schema     = "employees"
	TestTable1 = "dept_emp"
	TestTable2 = "employees"
)

func TestLookupTableMetadata(t *testing.T) {
	db, _ := GetDatabaseInstance(TEST_DB_DSN_STRING)
	defer db.Close()

	t.Run("Found", func(t *testing.T) {
		tableMap := NewTableMap(db)
		err := tableMap.Add(1, Schema, TestTable1)
		if err != nil {
			t.Fatal(err)
		}
		err = tableMap.Add(2, Schema, TestTable2)
		if err != nil {
			t.Fatal(err)
		}

		assertTableMetadata(t, &tableMap, 1, Schema, TestTable1)
		assertTableMetadata(t, &tableMap, 2, Schema, TestTable2)
		t.Log(tableMap)
	})

	t.Run("Fields", func(t *testing.T) {
		tableMap := NewTableMap(db)
		tableMap.Add(1, Schema, TestTable1)

		tableMetadata, ok := tableMap.LookupTableMetadata(1)
		t.Log(tableMetadata)

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
