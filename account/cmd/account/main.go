package main

import (
	"log"
	"time"

	"github.com/Sonam060703/microserviceGO/account"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
)

// Config has a DatabaseURL which it takes from envconfig
type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	// envconfig -> library used to read environment variables .
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r account.Repository
	// Keep retry to connect to db every 2 sec
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		// NewPostgresRepository take url return repository with db instance
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	// this will return an accountservice
	s := account.NewService(r)
	// Grpc server start listening at port 8080 with the accountservice on it .
	log.Fatal(account.ListenGRPC(s, 8080))
}
