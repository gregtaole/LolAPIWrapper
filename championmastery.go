package lolapiwrapper

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	masteryRootURL = "champion-mastery/v4/"
	bySummonerURL  = "champion-masteries/by-summoner"
	byChampionURL  = "by-champion"
	scoresURL      = "scores/by-summoner"
)

//ChampionMasteryDTO is the struct containing the information about a champion's mastery
type ChampionMasteryDTO struct {
	ChestGranted                 bool   `json:"chestGranted"`
	ChampionLevel                int    `json:"championLevel"`
	ChampionPoints               int    `json:"championPoints"`
	ChampionID                   int    `json:"championId"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	LastPlayTime                 int    `json:"lastPlayTime"`
	TokensEarned                 int    `json:"tokensEarned"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	SummonerID                   string `json:"summonerId"`
}

// ChampionMasteriesBySummoner gets all champion masteries for summonerID
func (c *client) ChampionMasteriesBySummoner(ctx context.Context, summonerID string) ([]ChampionMasteryDTO, error) {
	var res []ChampionMasteryDTO
	url := filepath.Join(masteryRootURL, bySummonerURL, summonerID)
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ChampionMasteriesBySummonerByChampion gets the mastery score for championID for summonerID
func (c *client) ChampionMasteriesBySummonerByChampion(ctx context.Context, summonerID string, championID int) (ChampionMasteryDTO, error) {
	var res ChampionMasteryDTO
	url := filepath.Join(masteryRootURL, bySummonerURL, summonerID, byChampionURL, strconv.Itoa(championID))
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// ScoresBySummoner gets the total mastery score for summonerID
func (c *client) ScoresBySummoner(ctx context.Context, summonerID string) (int, error) {
	var res int
	url := filepath.Join(masteryRootURL, scoresURL, summonerID)
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (cmd ChampionMasteryDTO) String() string {
	var str strings.Builder
	str.WriteString("ChampionMasteryDTO{\n")
	str.WriteString(fmt.Sprintf("    ChampionID: %v,\n", cmd.ChampionID))
	str.WriteString(fmt.Sprintf("    ChampionLevel: %v,\n", cmd.ChampionLevel))
	str.WriteString(fmt.Sprintf("    ChampionPoints: %v,\n", cmd.ChampionPoints))
	str.WriteString(fmt.Sprintf("    ChampionPointsSinceLastLevel: %v,\n", cmd.ChampionPointsSinceLastLevel))
	str.WriteString(fmt.Sprintf("    ChampionPointsUntilNextLevel: %v,\n", cmd.ChampionPointsUntilNextLevel))
	str.WriteString(fmt.Sprintf("    ChestGranted: %v,\n", cmd.ChestGranted))
	str.WriteString(fmt.Sprintf("    LastPlayTime: %v,\n", cmd.LastPlayTime))
	str.WriteString(fmt.Sprintf("    TokensEarned: %v,\n", cmd.TokensEarned))
	str.WriteString(fmt.Sprintf("    SummonerID: %v,\n", cmd.SummonerID))
	str.WriteString("},\n")
	return str.String()
}
