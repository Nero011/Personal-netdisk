package mysqlutil

import (
	"database/sql"
	"testing"
)

var testDb, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/netstore_user")

func Test_Init(t *testing.T) {
	i := 10
	for i >= 0 {
		db := Init()
		if db == nil {
			t.Fatal("db is nil")
		}
		err := db.Ping()
		if err != nil {
			t.Error("db conn error")
		}
		i--
	}

}
