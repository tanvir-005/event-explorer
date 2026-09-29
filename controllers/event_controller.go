package controllers

import (
	"github.com/beego/beego/v2/server/web"

	"github.com/tanvir-005/event-explorer/services"
)

type EventController struct {
	web.Controller
	EventService services.EventService
}

func (c *EventController) Get() {
	eventID := c.Ctx.Input.Param(":eventId")

	if eventID != "" {
		c.Data["EventID"] = eventID
		c.TplName = "details.tpl"
		return
	}

	city := c.GetString("city")
	countryCode := c.GetString("countryCode")

	c.Data["City"] = city
	c.Data["CountryCode"] = countryCode

	if city == "" || countryCode == "" {
		c.Data["Error"] = "Please select a city before searching."
		c.TplName = "listing.tpl"
		return
	}

	if c.EventService == nil {
		c.Data["Error"] = "Event service is not configured."
		c.TplName = "listing.tpl"
		return
	}

	music, sports, musicErr, sportsErr := c.EventService.GetEvents(
		city,
		countryCode,
	)

	c.Data["MusicEvents"] = music
	c.Data["SportsEvents"] = sports
	c.Data["MusicError"] = musicErr
	c.Data["SportsError"] = sportsErr

	c.TplName = "listing.tpl"
}
