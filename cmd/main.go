package main

import (
	"log"
	"os"
	testapi "testAPI"
	"testAPI/pkg/handlers"
	"testAPI/pkg/repository"
	"testAPI/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	db, err := repository.InitDb(repository.InitDbConfig())
	if err != nil {
		panic(err)
	}

	if len(os.Args) >= 2 {
		command := os.Args[1]
		switch command {
		case "migrate":
			if err := repository.RunMigration(db); err != nil {
				panic(err)
			}
			log.Println("Migration completed")
			return
		default:
			panic("Argument error")
		}
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handlers.NewHandler(service)
	server := testapi.NewServer()

	if err := server.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("%s", err)
	}
}
