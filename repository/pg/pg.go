package pg

import (
	"os"

	"github.com/jinzhu/gorm"
	// import specific dialect for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

// NewConnection creates new connection instance
func NewConnection() *gorm.DB {
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASS")
	dbname := os.Getenv("PG_DBNAME")
	dsnString := "host=" + host + " port=" + port + " user=" + user + " dbname=" + dbname + " password=" + pass + " sslmode=disable"
	log.Info("PG DSN:" + dsnString)
	db, err := gorm.Open("postgres", dsnString)

	if err != nil {
		panic("Could not connect to the Database:" + err.Error())
	}

	return db
}
