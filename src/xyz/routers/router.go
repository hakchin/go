package routers

import (
	"github.com/astaxie/beego"
	c "xyz/controllers"
)

func init() {
    beego.Router("/", &c.MainController{}, "get:Index")
}
