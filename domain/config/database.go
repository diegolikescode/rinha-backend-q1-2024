package config

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func SetupCassandra() {
    cluster := gocql.NewCluster("127.0.0.1")
    cluster.Port = 9042
    cluster.Keyspace = "rinha"

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "rinha",
		Password: "rinha",
	}

    cluster.Consistency = gocql.Quorum

	var err error

    Session, err = cluster.CreateSession()
    if err != nil {
	fmt.Println(err)
        panic("ERROR TRYING TO CONNECT TO CASSANDRA")
    }

	if err := Session.Query(
		"CREATE KEYSPACE IF NOT EXISTS rinha WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1}").Exec(); err != nil {
		panic("ERROR EXECUTING THE CREATION OF THE KEYSPACE")
	}

	log.Println("Cassandra WAS SUCCESSFULLY CONFIGURED")
}

