package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	otelgorm "gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

func Init() {
	dsn := "host=postgresql user=apicurio password=apicurio dbname=test_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
			panic("Failed to get sql.DB from gorm")
	}

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)

	_ = db.Use(otelgorm.NewPlugin())

	DB = db
}
