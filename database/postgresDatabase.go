package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/Arismonx/van-booking-system/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	*gorm.DB
}

var (
	postgreaDatabaseInstance *postgresDatabase
	once                     sync.Once
)

func NewPostgresDatabase(config *config.Database) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
			config.Host,
			config.Port,
			config.User,
			config.Password,
			config.DBName,
			config.SSLMode,
			config.Schema,
		)
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal("failed to connect database:", err)
		}

		log.Printf("Connected to database %s", config.DBName)

		postgreaDatabaseInstance = &postgresDatabase{conn}
	})

	return postgreaDatabaseInstance
}
func (d *postgresDatabase) Connection() *gorm.DB {
	return postgreaDatabaseInstance.DB
}
