package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lfcifuentes/clean-arquitecture/cmd"
	"github.com/spf13/viper"
)

func main() { // Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Configurar Viper para leer variables de entorno
	viper.AutomaticEnv()
	// Execute the root command
	cmd.Execute()
}
