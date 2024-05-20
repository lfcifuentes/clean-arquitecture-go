package cmd

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/fatih/color"
	"github.com/lfcifuentes/clean-arquitecture/pkg/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	// Add the serve command to the root command
	rootCmd.AddCommand(migrateCmd)
}

// migrateCmd represents the serve command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run the database migrations",
	Long:  `Run the database migrations to create the tables and the initial data`,
	Run: func(_ *cobra.Command, _ []string) {
		// Call the migrate function
		migrate()
	},
}

func migrate() {
	fmt.Println("Migrating the database...")
	color.Cyan("Connecting to PostgreSQL...")
	dbName := viper.GetString("DB_NAME")
	dns := db.GeneratePostgresConnStringWitoutDbName()

	// Intenta conectar a PostgreSQL sin seleccionar una base de datos específica
	db, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatal("Error connecting to PostgreSQL:", err)
	}
	err = db.Ping()

	if err != nil {
		log.Fatal("Error pining to PostgreSQL:", err)
	}

	defer db.Close()
	// Ejecuta la sentencia SQL para crear la base de datos si no existe
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s WITH ENCODING 'UTF8';", dbName))
	if err != nil {
		color.Red("Error creating database: %s /\n", err.Error())
	} else {
		color.Green("Database created or already exists.")
	}
	db.Close()
	dns = fmt.Sprintf("%s dbname=%s sslmode=disable", dns, dbName)
	db, err = sql.Open("postgres", dns)
	if err != nil {
		log.Fatal("Error connecting to the created database:", err)
	}
	defer db.Close()
	color.Cyan("Migrating the database...")

	createTableQuery := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
			username VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		color.Red("Failed to create users table: %s /\n", err.Error())
	}
	color.Green("Users table created successfully.")

	color.Cyan("Migration completed successfully")
}
