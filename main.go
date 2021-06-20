package main

import (
	"github.com/annabkr/paydayz/api"
	a "github.com/annabkr/paydayz/app"
)

func main() {
	app := a.Initialize()
	api.Initialize(app.GetRouter())
	app.Run()
}
