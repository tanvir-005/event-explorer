package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
	_ "github.com/tanvir-005/event-explorer/routers"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found; using environment variables")
	}
}

func main() {
	beego.Run()
}
