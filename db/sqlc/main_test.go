package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/StevenSopilidis/BackendMasterClass/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

// entry point to our tests
func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Could not load configuration: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Could not connect to db", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
