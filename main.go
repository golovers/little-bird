package main

import (
	"gitlab.com/7chip/little-bird/backend/api"
	"google.golang.org/appengine"
)

func main() {
	api.RegisterHandlers()
	appengine.Main()
}
