package app

import (
	"errors"
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/spf13/viper"
)

type Database struct {
	cfg     *viper.Viper
	cluster *gocql.ClusterConfig
}

func NewDatabase(cfg *viper.Viper) *Database {
	db := &Database{
		cfg: cfg,
	}

	db.cluster = gocql.NewCluster(cfg.GetString(CassandraHost))
	db.cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cfg.GetString(CassandraUser),
		Password: cfg.GetString(CassandraPass),
	}

	return db
}

func (db *Database) WaitFor() error {
	waitChannel := make(chan bool)
	timeout := db.cfg.GetInt(CassandraTimeout)

	go db.waitForCassandra(waitChannel)

	for i := 0; i < timeout; i++ {
		select {
		case <-waitChannel:
			return nil
		case <-time.After(time.Second):
			fmt.Printf(".")
		}
	}
	return errors.New("cassandra did not start in time")
}

func (db *Database) waitForCassandra(c chan bool) error {
	for {
		_, err := db.cluster.CreateSession()
		if err == nil {
			c <- true
			return nil
		}
		time.Sleep(time.Millisecond * 100)
	}
}
