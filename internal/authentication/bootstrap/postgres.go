package bootstrap

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration
	Builder      squirrel.StatementBuilderType
	Pool         *pgxpool.Pool
}

func NewPostgresDB(env *Env) *Postgres {
	pg := &Postgres{
		maxPoolSize:  env.DbMaxPoolSize,
		connAttempts: env.DbMaxAttempts,
		connTimeout:  time.Duration(env.DbTimeout),
	}

	ctx, cancel := context.WithTimeout(context.Background(), pg.connTimeout*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(GetPostgresConnectionString(env))
	if err != nil {
		log.Fatalf("Failed to parse connection string: %v\n", err)
	}
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Failed to initialize pool structure: %v\n", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Failed to physically connect: %v", err)
	} else {
		log.Println("Successfully connected to postgres database")
	}
	_, createTableErr := pool.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL
    )`)
	if createTableErr != nil {
		log.Fatalf("Failed to create table user %v", createTableErr)
	}
	return pg
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

func GetPostgresConnectionString(env *Env) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.DbUsername, env.DbPassword, env.Datasource, strconv.Itoa(env.DbPort), env.DbName)
}
