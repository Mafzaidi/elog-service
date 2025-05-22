package main

import (
	"log"

	"github.com/mafzaidi/elog/config"
	"github.com/mafzaidi/elog/internal/server"
	"github.com/mafzaidi/elog/pkg/db/mongodb"
)

func main() {
	cfg := config.GetConfig()
	db, err := mongodb.NewMongoDB(cfg)
	if err != nil {
		panic(err)
	}

	s := server.NewServer(cfg, db)
	if err := s.Run(); err != nil {
		log.Fatalf("Server could not be ran: %v", err)
	}
}
