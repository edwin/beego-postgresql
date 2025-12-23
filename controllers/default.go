package controllers

import (
        "beego-postgresql/database"
	"beego-postgresql/models"
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	var customers []models.Customer
        database.DB.Find(&customers) 
	
	c.Data["json"] = customers
	c.ServeJSON()
}
