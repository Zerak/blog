package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

const(
	WEB_SITE = "zerak.top"
	EMAIL = "Kingbug2010#126.com"
)

// @Title Get
// @Description get index
// @Success
// @router / [get]
func (c *MainController) Get() {
	c.Data["Website"] = WEB_SITE
	c.Data["Email"] = EMAIL
	c.TplName = "index.tpl"
}

// @Title Resume
// @Description get resume
// @Success
// @router /resume [get]
func (c *MainController) Resume() {
	c.Data["Website"] = WEB_SITE
	c.Data["Email"] = EMAIL
	c.TplName = "resume.tpl"
}