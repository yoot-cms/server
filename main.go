package main

import (
	"server/database"

	"github.com/joho/godotenv"
)

func init() {
  if err := godotenv.Load(); err!= nil {
    panic(err)
  }
}

func main() {
  db := database.GetDBConnectionPool()
  _ = db

}
