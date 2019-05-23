package lolapiwrapper

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
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
func (c *client) ChampionMasteriesBySummoner(ctx context.Context, summonerID string) (*[]ChampionMasteryDTO, error) {
	var res []ChampionMasteryDTO
	url := filepath.Join(masteryRootURL, bySummonerURL, summonerID)
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// ChampionMasteriesBySummonerByChampion gets the mastery score for championID for summonerID
func (c *client) ChampionMasteriesBySummonerByChampion(ctx context.Context, summonerID string, championID int) (*ChampionMasteryDTO, error) {
	var res ChampionMasteryDTO
	url := filepath.Join(masteryRootURL, bySummonerURL, summonerID, byChampionURL, strconv.Itoa(championID))
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// ScoresBySummoner gets all mastery scores for summonerID
func (c *client) ScoresBySummoner(ctx context.Context, summonerID string) (*int, error) {
	var res int
	url := filepath.Join(masteryRootURL, scoresURL, summonerID)
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (cmd ChampionMasteryDTO) String() string {
	return fmt.Sprintf("%v: %v\n", cmd.ChampionID, cmd.ChampionLevel)
}
