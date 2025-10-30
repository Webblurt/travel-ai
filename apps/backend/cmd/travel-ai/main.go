package main

import (
	"log"
	"net/http"
	"os"
	routes "travel-ai/internal/api/routes"
	clients "travel-ai/internal/clients"
	repositories "travel-ai/internal/repositories"
	services "travel-ai/internal/services"
	utils "travel-ai/internal/utils"

	"github.com/joho/godotenv"
)

func main() {

	// loading .env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found, using system environment variables", err)
	}

	//loading configuration
	cfg, err := utils.LoadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatal("Error loading config file", err)
	}

	// creating logger
	log := utils.NewLogger(cfg.Logger.LogLevel)

	// repository creation
	repo, err := repositories.NewRepository(cfg, log)
	if err != nil {
		log.Fatal("Error creating repository: ", err)
	}
	log.Info("Repository created successful")

	// start migrations
	if err := repo.RunMigrations(cfg); err != nil {
		log.Warn("Error running migrations: ", err)
	}
	log.Info("Migrations applied successfully")

	// creating clients for external apis
	clientsList, err := clients.CreateClients(cfg, log)
	if err != nil {
		log.Fatal("Error creating clients: ", err)
	}
	log.Info("Clients created successful")

	// creating service
	service, err := services.NewService(cfg, clientsList, repo, log)
	if err != nil {
		log.Fatal("Error creating service: ", err)
	}
	log.Info("Service created successful")

	// creating routes
	router, err := routes.CreateRoutes(service)
	if err != nil {
		log.Fatal("Error creating routes: ", err)
	}
	log.Info("Routes created successful")

	// starting http server
	log.Info("Starting the server on port ", cfg.Server.Port)
	if err := http.ListenAndServe(":"+cfg.Server.Port, router); err != nil {
		log.Fatal("Error starting server: ", err)
	}
	log.Info("Server started successful")
}
