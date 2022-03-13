package main

import (
	"fake_twitter/model"
	"fake_twitter/routes"
)

func main() {

	model.InitDb()

	routes.InitRouter()
}
