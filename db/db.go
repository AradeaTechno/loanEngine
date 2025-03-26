package db

import (
	"amarthaloan/config"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

func Init() {
	conf := config.DBConfig()

	// DSN set
	dsn := "host=" + conf.DB_HOST +
		" port=" + conf.DB_PORT +
		" user=" + conf.DB_USERNAME +
		" password=" + conf.DB_PASSWORD +
		" dbname=" + conf.DB_NAME +
		" sslmode=disable TimeZone=Asia/Jakarta"

	// Configure Logger
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	// Configure section pooling
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
}

func CreateConn() *gorm.DB {
	return db
}
