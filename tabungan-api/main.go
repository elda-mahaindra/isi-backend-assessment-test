package main

import (
	"database/sql"
	"fmt"
	"os"

	"tabungan-api/db/store/postgres_store"
	"tabungan-api/handler"
	"tabungan-api/service"
	"tabungan-api/utils/config"
	"tabungan-api/utils/errs"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	// init logger
	var logger = logrus.New()
	logger.Formatter = new(logrus.JSONFormatter)
	logger.Formatter = new(logrus.TextFormatter)                     //default
	logger.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	logger.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout

	// load environment variables from .env file
	config, err := config.LoadConfig(".")
	if err != nil {
		const op errs.Op = "config/LoadConfig"

		e := errs.E(op, errs.Database, fmt.Sprintf("failed to load config: %s", err))

		logger.WithFields(logrus.Fields{
			"op": op,
		}).Debug(e.Error())

		return
	}

	// create db connection
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		const op errs.Op = "sql/Open"

		e := errs.E(op, errs.Database, fmt.Sprintf("cannot connect to db: %s", err))

		logger.WithFields(logrus.Fields{
			"op": op,
		}).Debug(e.Error())

		return
	}

	// init data access layer
	store := postgres_store.NewPostgresStore(logger, conn)

	// init service layer
	service := service.NewService(logger, store)

	// init handler
	handler := handler.NewHandler(service)

	// init fiber app
	app := fiber.New()

	// endpoints
	app.Post("/daftar", handler.Registration)
	app.Post("/tabung", handler.Deposit)
	app.Post("/tarik", handler.Withdrawal)
	app.Get("/saldo/:no_rekening", handler.BalanceCheck)
	app.Get("/mutasi/:no_rekening", handler.TransactionHistory)

	// start the server
	host := config.Host
	port := config.Port

	err = app.Listen(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		const op errs.Op = "app/Listen"

		e := errs.E(op, errs.Database, fmt.Sprintf("failed to listen at port '%s': %s", port, err))

		logger.WithFields(logrus.Fields{
			"op": op,
		}).Debug(e.Error())

		return
	}
}
