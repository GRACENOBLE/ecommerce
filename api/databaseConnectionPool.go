package api

import "github.com/jackc/pgx/v5/pgxpool"

type DBConfig struct {
	DB *pgxpool.Pool
}