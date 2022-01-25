package main

import (
	_ "beegoproject/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
