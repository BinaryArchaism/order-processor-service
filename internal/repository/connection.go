package repository

import (
	"context"
	"database/sql"
	"embed"
	"github.com/BinaryArchaism/order-processor/internal/config"
	"github.com/rs/zerolog/log"
	"time"

	"github.com/pressly/goose/v3"

	_ "github.com/go-sql-driver/mysql"
)

type Connection interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

var _ Connection = new(TiDBConnection)

type TiDBConnection struct {
	conn *sql.DB
}

func (t TiDBConnection) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	queryContext, err := t.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return queryContext, nil
}

func Connect(ctx context.Context, cfg config.Config) (*TiDBConnection, error) {
	db, err := sql.Open("mysql", cfg.DBConfig.ConnectionString())
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(time.Minute * 3)

	go func() {
		<-ctx.Done()
		err := db.Close()
		if err != nil {
			log.Err(err).Msg("Error closing database connection")
			return
		}
		log.Info().Msg("Database connection closed")
	}()

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Info().Msg("Database connection established")

	tiDB := TiDBConnection{db}
	err = tiDB.InitMigrations()
	if err != nil {
		return nil, err
	}
	log.Info().Msg("Database migrations initialized")

	return &tiDB, nil
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func (t TiDBConnection) InitMigrations() error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	if err := goose.Up(t.conn, "migrations"); err != nil {
		return err
	}

	return nil
}
