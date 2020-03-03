package routers

import (
	"beegotest/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/gongluck", &controllers.GongluckController{})
}
