package models

type LeaderboardEntry struct {
	Nickname string `json:"nickname"`
	Record   int    `json:"record"`
}
