package postgres_store

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"tabungan-api/db/sqlc"
	"tabungan-api/utils/config"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var testQueries *sqlc.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	logger = logrus.New()
	logger.Formatter = new(logrus.JSONFormatter)
	logger.Formatter = new(logrus.TextFormatter)                     //default
	logger.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	logger.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout

	config, err := config.LoadConfig("../../..")
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %s", err)
	}

	testQueries = sqlc.New(testDB)

	os.Exit(m.Run())
}
