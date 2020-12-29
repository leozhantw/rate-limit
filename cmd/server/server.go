package main

import (
	"flag"
	"log"

	"github.com/leozhantw/rate-limit/internal/server"
)

var port = flag.String("port", ":3000", "Port to listen on")

func main() {
	srv := server.New().Start()

	if err := srv.Run(*port); err != nil {
		log.Fatalln(err)
	}
}
