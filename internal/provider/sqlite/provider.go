package sqlite

import (
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Singleon
var (
	once sync.Once
	db   *gorm.DB
	err  error
)

// Initialize a new connection with Database
func Init() *gorm.DB {
	// Execute just one time
	once.Do(func() {

		// Connect with database
		db, err = gorm.Open("sqlite3", os.Getenv("DB_FILE"))
		if err != nil {
			log.WithField("provider", "sqlite").Fatal(err)
		}

		// Define custom logger
		db.LogMode(true)
		db.SetLogger(&Logger{})
	})

	return db
}
