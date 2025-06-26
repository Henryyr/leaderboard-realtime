package main

import (
	"leaderboard-realtime/config"
	"leaderboard-realtime/handler"
	ws "leaderboard-realtime/websocket"

	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/websocket/v2"
)

func main() {
	_ = godotenv.Load(".env")

	app := fiber.New()

	appEnv := os.Getenv("APP_ENV")

	// CORS config: allow all origins in production, allow localhost in dev
	if appEnv == "production" {
		app.Use(cors.New())
	} else {
		app.Use(cors.New(cors.Config{
			AllowOrigins:     "http://localhost:3100,http://127.0.0.1:3100",
			AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
			AllowCredentials: true,
		}))
	}

	// Path public folder
	var publicPath string
	if appEnv == "production" {
		publicPath = "./public"
	} else {
		publicPath = "../public"
	}

	app.Use("/", filesystem.New(filesystem.Config{
		Root:   http.Dir(publicPath),
		Browse: true,
		Index:  "client.html",
		MaxAge: 3600,
	}))

	config.ConnectRedis()

	app.Post("/submit", handler.SubmitScore)
	app.Get("/leaderboard", handler.GetLeaderBoard)
	app.Get("/ws", ws.WebSocketHandler, websocket.New(ws.WebSocketConn))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3100"
	}
	log.Fatal(app.Listen(":" + port))
}
