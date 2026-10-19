package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB(env *Env) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(env.contextTimeout)*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(env.GetPostgresConnectionString())
	if err != nil {
		log.Fatalf("Failed to parse connection string: %v\n", err)
	}
	pool, err := pgxpool.NewWithConfig(ctx, config)
	log.Println("Successfully connected to PostgreSQL natively using pgxpool!")
	return pool
}

func (env *Env) GetPostgresConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.dbUsername, env.dbPassword, env.datasource, env.dbPort, env.dbName)
}
