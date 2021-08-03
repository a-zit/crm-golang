package main

import (
	"crm-golang/routers"
)

func main() {
	e := routers.SetupRouter()
	e.Logger.Fatal(e.Start(":8080"))
}
