package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type GongluckController struct {
	beego.Controller
}

func (c *GongluckController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "1039994845@qq.com"
	c.TplName = "index.tpl"
}
