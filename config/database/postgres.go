package database

import (
	"context"
	"fmt"
	"log"

	"github.com/naelcodes/ab-backend/internal/configs"
	"github.com/naelcodes/ab-backend/internal/ent"
	"github.com/naelcodes/ab-backend/internal/pkg/logger"

	_ "github.com/lib/pq"
)

func PostgresConnection(context context.Context, logger *logger.Logger) *ent.Client {

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", configs.POSTGRES_DB_HOST, configs.POSTGRES_DB_PORT, configs.POSTGRES_DB_USER, configs.POSTGRES_DB_NAME, configs.POSTGRES_DB_PASSWORD)

	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	logger.Info("Database Connected ....")

	return client
}
