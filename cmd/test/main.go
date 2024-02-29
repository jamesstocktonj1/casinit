package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocql/gocql"
)

func main() {
	cluster := gocql.NewCluster(os.Getenv("CASSANDRA_HOST"))
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: os.Getenv("CASSANDRA_USER"),
		Password: os.Getenv("CASSANDRA_PASS"),
	}

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	fmt.Println("Cassandra is available!")
}
