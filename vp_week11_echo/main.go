package main

import (
	"vp_week11_echo/db"
	"vp_week11_echo/routes"
)
func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":9090"))
}
