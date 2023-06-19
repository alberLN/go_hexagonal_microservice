package main

import (
	"example/microservice/internal/repository"
	"example/microservice/internal/service"
	"example/microservice/internal/transport/bbdd"
	"example/microservice/internal/transport/http"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../internal/infraestructure/environment/.env")
	if err != nil {
		log.Println(err)
		return
	}
	uri := os.Getenv("POSTGRESS_URI")
	maxIdleConnsStr := os.Getenv("POSTGRESS_MAX_IDLE_CONNS")
	maxIdleConns, err := strconv.Atoi(maxIdleConnsStr)
	if err != nil {
		log.Printf("Cast Error: %s", err)
		return
	}
	maxOpenConnsString := os.Getenv("POSTGRESS_MAX_OPEN_CONNS")
	maxOpenConns, err := strconv.Atoi(maxOpenConnsString)
	if err != nil {
		log.Printf("Cast Error: %s", err)
		return
	}
	port := os.Getenv("API_PORT")

	db, err := bbdd.NewBBDD(uri, maxIdleConns, maxOpenConns)
	if err != nil {
		log.Println(err)
		return
	}

	taskRepo := repository.NewTaskRepositoryImpl(db)
	taskService := service.NewTaskServiceImpl(taskRepo)

	router := gin.Default()
	http.RegisterHandlers(router, taskService)

	router.Run(port)
}
