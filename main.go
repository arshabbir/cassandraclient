package main

import (
	"log"
	"time"

	"github.com/arshabbir/cassandraclient/domain/dao"
	"github.com/arshabbir/cassandraclient/domain/dto"
)

func main() {
	//Create the DAO object

	dao := dao.NewDAO()

	if dao == nil {
		log.Println("DAO creation error")
		return
	}

	log.Println("DAO creation successful .")

	if err := dao.Create(dto.Student{Id: 7, Name: "arshabbir", Marks: 90, Class: "Phd"}); err != nil {

		log.Println("Error inserting into dao")
		return
	}

	log.Println("Insertion  successful .")

	log.Println("Sleeping for 10 seconds......")
	time.Sleep(10 * time.Second)

	log.Println("Updating the record.......")

	if err := dao.Update(7, dto.Student{Id: 7, Name: "arshabbirhussain", Marks: 93, Class: "(Phd)"}); err != nil {

		log.Println("Error updating  into dao")
		return
	}

	log.Println("Updation  successful .")

	log.Println("Sleeping for 60 seconds......")
	time.Sleep(60 * time.Second)

	log.Println("Deleting the record.......")

	if err := dao.Delete(7); err != nil {

		log.Println("Error deleteting into dao")
		return
	}

	log.Println("Deletion  successful .")

}
