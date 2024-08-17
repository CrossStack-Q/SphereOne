package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/CrossStack-Q/SphereOne/backend/api"
	"github.com/CrossStack-Q/SphereOne/backend/storage"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "server address")

	flag.Parse()

	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, store)

	fmt.Println("Server Started ", *listenAddr)

	log.Fatal(server.Start())
}
