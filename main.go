package main

import (
	"framework_seed/routes"
)

func main() {
	r := routes.InitRoutes()
	r.Run(":8765")
}
