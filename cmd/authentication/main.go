package main

import (
	"chatAppGo/internal/authentication/bootstrap"

	"log"
)

func main() {
	log.Println("Starting Authentication service...")
	env := bootstrap.NewEnv()
	bootstrap.NewPostgresDB(env)
}
