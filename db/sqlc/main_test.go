package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/RickyJia2018/simple_bank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := util.LaodConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to databse")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
