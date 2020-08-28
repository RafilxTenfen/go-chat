package database

import (
	"fmt"
	"time"

	"github.com/RafilxTenfen/go-chat/app"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rhizomplatform/log"
)

// OpenDBConnection open the database connection using GORM
func OpenDBConnection(s Settings) (*gorm.DB, error) {
	for i := 0; i < 5; i++ {
		if i > 0 {
			time.Sleep(time.Duration(i+1) * time.Second)
		}

		log.With(log.F{
			"tries": i + 1,
		}).Info("Connecting to database")
		db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", s.Host, s.Port, s.User, s.DatabaseName, s.Password))
		if err != nil {
			log.Error(err)
			continue
		}

		err = db.DB().Ping()
		if err != nil {
			log.Error(err)
		} else {
			autoMigrateStructs(db)
			return db, nil
		}
	}

	return nil, fmt.Errorf("error on connect to the database")
}

func autoMigrateStructs(db *gorm.DB) {
	db.AutoMigrate(&app.User{})
}

// DBConnect opens the database connect loading the settings from the .env file
func DBConnect() (*gorm.DB, error) {
	st, err := SettingsFromEnv()
	if err != nil {
		return nil, err
	}

	return OpenDBConnection(st)
}
