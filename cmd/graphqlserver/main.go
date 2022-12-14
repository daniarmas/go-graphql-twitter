package main

import (
	"context"
	"fmt"
	"log"

	"github.com/daniarmas/gographqltwitter/config"
	"github.com/daniarmas/gographqltwitter/postgres"
)

func main() {
	ctx := context.Background()

	config := config.New()

	db := postgres.New(ctx, config)

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server running")
}
