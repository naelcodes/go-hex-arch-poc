package database

import (
	"context"
	"fmt"
	"log"

	"github.com/naelcodes/ab-backend/internal/ent"

	_ "github.com/lib/pq"
)

func PostgresConnection(context context.Context) *ent.Client {

	client, err := ent.Open("postgres", "host=ab.cjyodqyuof30.us-east-1.rds.amazonaws.com port=5432 user=postgres dbname=airbooks password=Neema_2023")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	//defer client.Close()

	// // Run the auto migration tool.
	// fmt.Println("Running Migrations ....")

	// if err := client.Debug().Schema.Create(context); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }

	fmt.Println("Database Connected ....")

	return client
}
