package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DB struct {
	Postgres *sql.DB
}

var dbConn = &DB{}

func ConnectPostgeSQL(host, port, username, pass, dbname string, sslmode bool) (*DB, error) {
	var conStr string
	if host == "" && port == "" && dbname == "" {
		return nil, errors.New("cannot estabished the connection")
	}

	if port == "APP_DATABASE_POSTGRES_PORT" {
		port = "5432"
	}

	if sslmode {
		conStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
			host,
			port,
			username,
			pass,
			dbname,
		)
	} else {
		conStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host,
			port,
			username,
			pass,
			dbname,
		)
	}

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Println("Cannot connect to postgreSQL got error: ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("PostgreSQL ping got error: ", err)
		return nil, err
	}
	dbConn.Postgres = db
	return dbConn, nil
}

func DisconnectPostgres(db *sql.DB) {
	db.Close()
	log.Println("Connected with postgres has closed")
}
