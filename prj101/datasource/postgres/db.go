package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	postgres_host     = "db"
	postgres_port     = 5432
	postgres_user     = "postgres"
	postgres_password = "postgres"
	postgres_dbname   = "hospitalregistration"
)

var (
	Client *sql.DB
)

func init() {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname)
	var err error
	if os.Getenv("DATABASE_URL") == "" {
		// getenv = get environment (trying to get the variable from the url )(it is from os package)
		Client, err = sql.Open("postgres", dataSourceName)
		// name of the driver (what we are using i.e postgress sql)
	} else {
		Client, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
		// anything that is fetch up from the os.Getenv("DATABASE_URL") will stored in the client
		// if there is erro then it will be stored in the err
	}
	if err != nil {
		panic(err)
	} else {
		log.Println("Database successfully configures")
	}
}
