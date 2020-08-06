package cassandra

import (
	"fmt"
	"log"

	"github.com/arshabbir/cassandraclient/domain/dto"
	"github.com/arshabbir/cassandraclient/utils"
	"github.com/gocql/gocql"
)

type client struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}
type Client interface {
	Create(dto.Student) *utils.ApiError
	Read(int) ([]dto.Student, *utils.ApiError)
	Update(int, dto.Student) *utils.ApiError
	Delete(int) *utils.ApiError
}

////

func init() {

	//Initialize Cassandra session
	/*cluster := gocql.NewCluster("34.207.221.180")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum*/

	var err error

	cluster := gocql.NewCluster("54.90.12.233")
	cluster.Keyspace = "student"
	cluster.ProtoVersion = 4
	cluster.DisableInitialHostLookup = true
	cluster.Timeout = 50000
	cluster.Port = 9042

	_, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")

}

////

func NewDBClient() Client {
	//Get the Environment variable "CASSANDRACLUSTER"

	/*
		cluster := gocql.NewCluster("54.90.12.233")
		cluster.Keyspace = "student"

		session, err := cluster.CreateSession()

		if err != nil {
			log.Println("Cassandra Session Creation Error..", err.Error())
			return nil
		}

		//defer session.Close()
		return &client{cluster: cluster, session: session}*/
	return nil
}

func (c *client) Create(st dto.Student) *utils.ApiError {

	//Form the insert query & execute it

	insertQuery := fmt.Sprintf("INSERT INTO student(id, name, class, marks) values(?, ?, ?, ?)")

	if err := c.session.Query(insertQuery, st.Id, st.Name, st.Class, st.Marks).Exec(); err != nil {
		log.Println("Insert query error")
		return &utils.ApiError{Status: 0, Message: "Insert query error"}
	}

	return nil
}

func (c *client) Read(id int) ([]dto.Student, *utils.ApiError) {

	return nil, nil
}

func (c *client) Update(id int, st dto.Student) *utils.ApiError {

	updateQuery := fmt.Sprintf("Update student set  name=?, class=?, marks=? where id=?")

	if err := c.session.Query(updateQuery, st.Name, st.Class, st.Marks, id).Exec(); err != nil {
		log.Println("Update query error")
		return &utils.ApiError{Status: 0, Message: "Update query error"}
	}

	return nil
}

func (c *client) Delete(id int) *utils.ApiError {

	deleteQuery := fmt.Sprintf("Delete student  where id=?")

	if err := c.session.Query(deleteQuery, id).Exec(); err != nil {
		log.Println("Delete query error")
		return &utils.ApiError{Status: 0, Message: "Delete query error"}
	}
	return nil
}
