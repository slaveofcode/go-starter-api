package pg

import (
	"os"

	"github.com/jinzhu/gorm"
	// import specific dialect for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewConnection creates new connection instance
func NewConnection() *gorm.DB {
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASS")
	dbname := os.Getenv("PG_DBNAME")
	db, err := gorm.Open("postgres", "host="+host+" port="+port+" user="+user+" dbname="+dbname+" password="+pass+" sslmode=disable")

	if err != nil {
		panic("Couldn't connect to the database")
	}

	return db
}
