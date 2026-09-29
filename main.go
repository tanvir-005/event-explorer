package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/tanvir-005/event-explorer/routers"
)

func main() {
	beego.Run()
}
