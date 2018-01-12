package database

import (
	"testing"
	"fmt"
)

const TEST_DB_DSN_STRING = "root@tcp(0.0.0.0:3306)/employees"

func TestGetDatabaseInstance(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		db, err := GetDatabaseInstance(TEST_DB_DSN_STRING)
		if err != nil {
			t.Error(err)
		}
		t.Log(db)

		t.Run("Execute", func(t *testing.T) {
			var column string
			query := fmt.Sprintf(GetColumnsQuery, "employees", "dept_emp")
			err := db.QueryRow(query).Scan(&column)

			if err != nil {
				t.Error(err)
			}

			if column == "" {
				t.Fatal("Not Get Value From Database")
			}
			t.Log(column)
		})

		defer db.Close()
	})

	t.Run("Not Found", func(t *testing.T) {
		_, err := GetDatabaseInstance("root@/test")
		if err == nil {
			t.Log(err)
		}
	})
}