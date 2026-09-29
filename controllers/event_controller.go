package controllers

import "github.com/beego/beego/v2/server/web"

type EventController struct {
	web.Controller
}

func (c *EventController) Get() {
	eventID := c.Ctx.Input.Param(":eventId")

	if eventID != "" {
		c.Data["EventID"] = eventID
		c.TplName = "details.tpl"
		return
	}

	c.Data["City"] = c.GetString("city")
	c.Data["CountryCode"] = c.GetString("countryCode")
	c.TplName = "listing.tpl"
}
