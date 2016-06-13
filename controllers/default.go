package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// @Title Get
// @Description get index
// @Success
// @router / [get]
func (c *MainController) Get() {
	c.Data["Website"] = "zerak.me"
	c.Data["Email"] = "Kingbug2010@126.com"
	c.TplName = "index.tpl"
}

// @Title Resume
// @Description get resume
// @Success
// @router /resume [get]
func (c *MainController) Resume() {
	c.Data["Website"] = "zerak.me"
	c.Data["Email"] = "Kingbug2010@126.com"
	c.TplName = "resume.tpl"
}