package test

import (
	utils "Challenge/test/utils"
	"testing"
	"time"
)

func Test_Validate_CockroachDB_Success_Connection(t *testing.T) {
	_, err := utils.GetDatabaseConnection(&utils.Configs)
	if err != nil {
		t.Errorf("The test failed during connecting to database: %s", err)
	}
}

func Test_Create_Single_Query_For_Check_Version_Request_To_Database(t *testing.T) {
	var now time.Time
	db, err := utils.GetDatabaseConnection(&utils.Configs)
	if err != nil {
		t.Errorf("The test failed during connecting to database: %s", err)
	}
	db.Raw("SELECT NOW()").Scan(&now)
	if now.IsZero() {
		t.Errorf("The test failed during executing the query: %s", err)
	}

}
