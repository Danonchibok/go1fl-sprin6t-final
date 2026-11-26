package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	loger := log.New(os.Stdout, "server: ", log.LstdFlags)

	server := server.NewServer(loger)
	server.StartServer()
}
