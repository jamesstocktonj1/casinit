package main

import (
	"fmt"
	"log"

	"github.com/jamesstocktonj1/casinit/app"
	"github.com/spf13/viper"
)

func main() {
	cfg := viper.New()
	app.BindEnv(cfg)

	db := app.NewDatabase(cfg)

	fmt.Println("Waiting for cassandra to start")
	err := db.WaitFor()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nCassandra started")
}
