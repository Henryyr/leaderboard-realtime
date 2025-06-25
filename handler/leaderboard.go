package handler

import (
	"log"
	"strconv"

	"leaderboard-realtime/config"
	"leaderboard-realtime/model"
	ws "leaderboard-realtime/websocket"

	"github.com/gofiber/fiber/v2"
)

func SubmitScore(c *fiber.Ctx) error {
	var body model.SubmitScoreRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).SendString("Invalid Request")
	}

	// Ambil point, fallback ke score jika point kosong
	point := body.Point
	if point == 0 {
		point = body.Score
	}
	if point <= 0 {
		point = 1
	}

	log.Printf("DEBUG: SubmitScore received: username=%s, point=%v\n", body.Username, point)

	_, err := config.Rdb.ZIncrBy(config.Ctx, "leaderboard", point, body.Username).Result()
	if err != nil {
		return c.Status(500).SendString("Gagal Simpan Score")
	}

	ws.HubInstance.Broadcast([]byte("update"))

	return c.SendString("")
}

func GetLeaderBoard(c *fiber.Ctx) error {
	results, err := config.Rdb.ZRevRangeWithScores(config.Ctx, "leaderboard", 0, 9).Result()
	if err != nil {
		return c.Status(500).SendString("Gagal Ambil Leaderboard")
	}

	var leaderboard []map[string]string
	for i, item := range results {
		entry := map[string]string{
			"rank":     strconv.Itoa(i + 1),
			"username": item.Member.(string),
			"score":    strconv.FormatFloat(item.Score, 'f', 0, 64),
		}
		leaderboard = append(leaderboard, entry)
	}

	return c.JSON(leaderboard)
}
