package main

import (
	"manage_sales/configuration"
)

func main() {
	configuration.LoadEnv()
	configuration.ConnectDatabase()
	configuration.CreateHttp()
	
}