package routers

import (
	"github.com/beego/beego/v2/server/web"

	"github.com/tanvir-005/event-explorer/controllers"
)

func init() {
	web.Router("/", &controllers.HomeController{})
	web.Router("/events", &controllers.EventController{})
	web.Router("/events/:eventId", &controllers.EventController{})
}
