package main

import (
	"log"

	"github.com/arshabbir/cassandraclient/app"
)

func main() {

	//os.Setenv("CLUSTERIP", "3.90.70.181")
	//os.Setenv("PORT", ":8080")
	log.Println("Starting the Application.......")
	app.StartApplication()

	return
}
