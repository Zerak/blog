package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "zerak.me"
	c.Data["Email"] = "Kingbug2010@gmail.com"
	c.TplName = "index.tpl"
}
