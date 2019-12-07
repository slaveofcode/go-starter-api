package pg

import (
	"github.com/jinzhu/gorm"
	// import specific dialect for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

// Connection info
type Connection struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSL      bool
}

// NewConnection creates new connection instance
func NewConnection(conn *Connection) *gorm.DB {
	sslmode := "sslmode=disable"
	if conn.SSL {
		sslmode = "sslmode=enable"
	}

	dsnString := "host=" + conn.Host + " port=" + conn.Port + " user=" + conn.Username + " dbname=" + conn.DBName + " password=" + conn.Password + " " + sslmode
	log.Info("PG DSN:" + dsnString)
	db, err := gorm.Open("postgres", dsnString)

	if err != nil {
		panic("Could not connect to the Database:" + err.Error())
	}

	return db
}
