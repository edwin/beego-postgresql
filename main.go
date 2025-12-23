package main

import (
        "github.com/beego/beego/v2/server/web"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
        _ "beego-postgresql/routers"
)

var DB *gorm.DB

func init() {
	// Database connection string
	dsn := "host=postgresql user=apicurio password=apicurio dbname=test_db port=5432 sslmode=disable"
	
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
}

func main() {
        web.Run()
}

type MainController struct {
        web.Controller
}
