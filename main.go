package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fasthttp/session"
	"github.com/fasthttp/session/redis"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/slaveofcode/go-starter-api/context"
	"github.com/slaveofcode/go-starter-api/logger"
	"github.com/slaveofcode/go-starter-api/middleware"
	"github.com/slaveofcode/go-starter-api/repository/pg"
	"github.com/slaveofcode/go-starter-api/route"
	"github.com/valyala/fasthttp"
)

// sess handles the user session
var sess = session.New(session.NewDefaultConfig())

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	logger.Setup()

	// Session
	redisPort, _ := strconv.ParseInt(os.Getenv("REDIS_PORT"), 10, 64)
	sessCfg := &redis.Config{
		Host:        os.Getenv("REDIS_HOST"),
		Port:        redisPort,
		PoolSize:    8,
		IdleTimeout: 300,
		KeyPrefix:   "sess",
	}
	err = sess.SetProvider("redis", sessCfg)
	if err != nil {
		fmt.Println("Unable to intialize sesion with redis: ", err.Error())
	}
	fmt.Println("Session initialized...")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		panic("Please define port number!")
	}

	// Database initialization
	// host := os.Getenv("PG_HOST")
	// port := os.Getenv("PG_PORT")
	// user := os.Getenv("PG_USER")
	// pass := os.Getenv("PG_PASS")
	// dbname := os.Getenv("PG_DBNAME")
	db := pg.NewConnection(&pg.Connection{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		Username: os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASS"),
		DBName:   os.Getenv("PG_DBNAME"),
	})
	defer db.Close()

	svr := &fasthttp.Server{
		Handler: middleware.CORS(route.New(&context.AppContext{
			DB:      db,
			Session: sess,
		}).Handler),
		LogAllErrors: true,
		Logger:       &logrus.Logger{},
	}

	logrus.Info("Server Running on http://localhost:" + port)

	if err := svr.ListenAndServe(":" + port); err != nil {
		logrus.Fatalf("error in ListenAndServe: %s", err)
	}
}
