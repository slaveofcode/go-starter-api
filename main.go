package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/slaveofcode/go-starter-api/logger"
	"github.com/slaveofcode/go-starter-api/middleware"
	"github.com/slaveofcode/go-starter-api/route"
	"github.com/valyala/fasthttp"
)

var log = logger.New()

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	loadEnv()
	port := os.Getenv("PORT")

	if port == "" {
		panic("Please define port number!")
	}

	svr := &fasthttp.Server{
		Handler:      middleware.CORS(route.New().Handler),
		LogAllErrors: true,
		Logger:       log,
	}

	if err := svr.ListenAndServe(":" + port); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
