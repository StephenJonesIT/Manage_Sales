package main

import (
	"manage_sales/configuration"
)

func main() {
	configuration.ConnectDatabase()
	configuration.CreateHttp()
	
}