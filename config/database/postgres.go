package database

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/naelcodes/ab-backend/config"
	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/pkg/logger"
)

func PostgresConnection(context context.Context, logger *logger.Logger) *ent.Client {

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", config.POSTGRES_DB_HOST, config.POSTGRES_DB_PORT, config.POSTGRES_DB_USER, config.POSTGRES_DB_NAME, config.POSTGRES_DB_PASSWORD)

	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	logger.Info("Database Connected ....")

	return client
}
