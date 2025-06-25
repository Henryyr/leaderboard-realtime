package model

type SubmitScoreRequest struct {
	Username string  `json:"username"`
	Point    float64 `json:"point"`
	Score    float64 `json:"score"`
}
