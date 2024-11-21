package main

import (
	"log"
	"os"

	"github.com/danielmalka/challenge-bravo/config"
	"github.com/danielmalka/challenge-bravo/internal/flags"
	"github.com/danielmalka/challenge-bravo/internal/http/gin"
	"github.com/danielmalka/challenge-bravo/internal/server"
)

func main() {
	c, err := config.ReadConfig()
	if err != nil {
		log.Println(err)
		return
	}
	log.SetOutput(os.Stdout)
	flags.CheckFlags(c)

	log.Println("# API Initialized")

	initAPI(c)
}

func initAPI(c config.Config) {
	h := gin.Handlers(c)
	err := server.Start(c.Port, h)
	if err != nil {
		log.Fatalf("error running server: %s", err)
	}
}