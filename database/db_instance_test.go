package database

import (
	"testing"
)

const TEST_DB_DSN_STRING = "root@/employees"

func TestGetDatabaseInstance(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		db, err := GetDatabaseInstance(TEST_DB_DSN_STRING)
		if err != nil {
			t.Error(err)
		}
		t.Log(db)
	})

	t.Run("Not Found", func(t *testing.T) {
		_, err := GetDatabaseInstance("root@/test")
		if err == nil {
			t.Log(err)
		}
	})
}