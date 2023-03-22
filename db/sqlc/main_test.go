package db

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testDb *sql.DB
var testQuery Querier

const dbUrl = "postgresql://root:password@localhost:5431/go_payments_test?sslmode=disable"

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open("postgres", dbUrl)

	if err != nil {
		panic(err)
	}

	testQuery = NewStore(testDb)
	os.Exit(m.Run())
}
