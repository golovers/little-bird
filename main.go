package main

import (
	"gitlab.com/koffee/little-bird/backend/api"
	"google.golang.org/appengine"
)

func main() {
	api.RegisterHandlers()
	appengine.Main()
}
