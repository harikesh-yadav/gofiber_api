package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/harikesh-yadav/gofiber_api/database/queries"
	"github.com/harikesh-yadav/gofiber_api/utils"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Queries struct {
	*queries.UserQueries
}

func Connection() (*Queries, error) {

	var db *sqlx.DB
	var err error

	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	dbUrl, err := utils.ConnectionUrlBuilder("postgres")

	if err != nil {
		return nil, err
	}

	db, err = sqlx.Connect("pgx", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return &Queries{
		UserQueries: &queries.UserQueries{DB: db}, // from User model
	}, nil
}
