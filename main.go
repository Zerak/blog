package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run("0.0.0.0")
}

