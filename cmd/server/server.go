package main

import (
	"flag"

	"github.com/leozhantw/rate-limit/internal/server"
)

var port = flag.String("port", ":3000", "Port to listen on")

func main() {
	srv := server.New().Start()

	srv.Run(*port)
}
