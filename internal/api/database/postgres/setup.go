package postgres

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gitlab.com/revolutionize1/foward-api/internal/app"
)

var Instance *sqlx.DB

func dsn() string {
	return app.Instance.Config.Postgres.DSN
}

func open() error {
	database, err := sqlx.Connect("postgres", dsn())

	if err != nil {
		return err
	}

	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(5)

	if err := database.Ping(); err != nil {
		return err
	}

	Instance = database
	return nil
}

func Init() {
	if err := open(); err != nil {
		log.Fatal(err)
	}
	runMigrations()
}
