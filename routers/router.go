package routers

import (
	"github.com/tanvir-005/event-explorer/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
