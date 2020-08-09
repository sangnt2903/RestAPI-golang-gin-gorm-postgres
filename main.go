package main

import (
	"pipeline_security/models"
	"pipeline_security/routes"
)

func main(){
	r := routes.SetRoutesConfig()
	models.ConnectDataBase()

	r.Run()
}