package mysqlutil

import (
	"testing"
)

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
