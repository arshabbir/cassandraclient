package main

import (
	"log"

	"github.com/arshabbir/cassandraclient/domain/dao"
	"github.com/arshabbir/cassandraclient/domain/dto"
)

func main1() {
	//Create the DAO object

	dao := dao.NewDAO()

	if dao == nil {
		log.Println("DAO creation error")
		return
	}

	log.Println("DAO creation successful .")

	if err := dao.Create(dto.Student{Id: 1, Name: "arshabbir", Marks: 90, Class: "Phd"}); err != nil {

		log.Println("Error inserting into dao")
		return
	}

	log.Println("Insertion  successful .")
}
