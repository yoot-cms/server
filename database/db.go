package database

import "github.com/jmoiron/sqlx"

func GetDBConnectionPool(dbURL string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	return db
}
