package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fasthttp/session"
	"github.com/fasthttp/session/redis"
	"github.com/joho/godotenv"
	"github.com/slaveofcode/go-starter-api/context"
	"github.com/slaveofcode/go-starter-api/logger"
	"github.com/slaveofcode/go-starter-api/middleware"
	"github.com/slaveofcode/go-starter-api/repository/pg"
	"github.com/slaveofcode/go-starter-api/route"
	"github.com/valyala/fasthttp"
)

var log = logger.New()

// sess handles the user session
var sess = session.New(session.NewDefaultConfig())

func init() {
	loadEnv()
	// Session
	redisPort, _ := strconv.ParseInt(os.Getenv("REDIS_PORT"), 10, 64)
	sessCfg := &redis.Config{
		Host:        os.Getenv("REDIS_HOST"),
		Port:        redisPort,
		PoolSize:    8,
		IdleTimeout: 300,
		KeyPrefix:   "sess",
	}
	err := sess.SetProvider("redis", sessCfg)
	if err != nil {
		fmt.Println("Unable to intialize sesion with redis: ", err.Error())
	}
	fmt.Println("Session initialized...")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		panic("Please define port number!")
	}

	// Database initialization
	db := pg.NewConnection()
	defer db.Close()

	svr := &fasthttp.Server{
		Handler: middleware.CORS(route.New(&context.AppContext{
			DB:       db,
			Sesssion: sess,
		}).Handler),
		LogAllErrors: true,
		Logger:       log,
	}

	log.Info("Server Running on http://localhost:" + port)

	if err := svr.ListenAndServe(":" + port); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
