package schemas

import "encoding/json"

type MatchMetadata struct {
	DataVersion  string   `json:"dataVersion"`
	MatchId      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

type MatchInfo struct {
	GameCreation       int               `json:"gameCreation"`
	GameDuration       int               `json:"gameDuration"`
	GameStartTimestamp int               `json:"gameStartTimestamp"`
	GameEndTimestamp   int               `json:"gameEndTimestamp"`
	GameId             int               `json:"gameId"`
	GameMode           string            `json:"gameMode"`
	Participants       []json.RawMessage `json:"participants"`
}

type MatchInfoByMatchId struct {
	Metadata MatchMetadata `json:"metadata"`
	Info     MatchInfo     `json:"info"`
}
