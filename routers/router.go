package routers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"

	"github.com/tanvir-005/event-explorer/controllers"
	"github.com/tanvir-005/event-explorer/services"
)

func init() {
	ticketmasterService := services.NewTicketmasterService(http.DefaultClient)

	eventCache := services.NewMemoryEventCache()

	eventService := services.NewEventService(
		ticketmasterService,
		eventCache,
	)

	eventController := &controllers.EventController{
		EventService: eventService,
	}

	web.Router("/", &controllers.HomeController{})
	web.Router("/events", eventController)
	web.Router("/events/:eventId", eventController)
}
