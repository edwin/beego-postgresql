package main

import (
        "beego-postgresql/database"
	_ "beego-postgresql/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	database.Init()
	web.Run()
}

