package routers

import (
    "beego-postgresql/controllers"
    "github.com/beego/beego/v2/server/web"
)

func init() {
    // This line maps the root URL "/" to your MainController
    web.Router("/", &controllers.MainController{})
}
