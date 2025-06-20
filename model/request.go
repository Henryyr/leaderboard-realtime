package model

type SubmitScoreRequest struct {
	Username string  `json:"username"`
	Score    float64 `json:"score"`
}
