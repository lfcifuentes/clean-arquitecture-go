package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

func GeneratePostgresConnStringWitoutDbName() string {
	// DB credentials
	user := viper.GetString("DB_USER")
	pass := viper.GetString("DB_PASSWORD")
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")

	// Crear la cadena de conexión para PostgreSQL sin especificar la base de datos
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s sslmode=disable",
		user, pass, host, port,
	)
}
func GeneratePostgresConnString() string {
	// DB credentials
	dbname := viper.GetString("DB_NAME")

	return fmt.Sprintf("%s dbname=%s", GeneratePostgresConnStringWitoutDbName(), dbname)
}

func NewPostgresConn(dsn string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
