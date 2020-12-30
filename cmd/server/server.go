package main

import (
	"log"
	"os"

	"github.com/leozhantw/rate-limit/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	srv := server.New().Start()

	if err := srv.Run(":" + port); err != nil {
		log.Fatalln(err)
	}
}
