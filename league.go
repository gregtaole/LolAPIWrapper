package lolapiwrapper

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
)

// QueueType …
type QueueType string

//…
const (
	Solo   QueueType = "RANKED_SOLO_5x5"
	FlexSR QueueType = "RANKED_FLEX_SR"
	FlexTT QueueType = "RANKED_FLEX_TT"
)

// Tier …
type Tier string

//
const (
	Iron     Tier = "IRON"
	Bronze   Tier = "BRONZE"
	Silver   Tier = "SILVER"
	Gold     Tier = "GOLD"
	Platinum Tier = "PLATINUM"
	Diamond  Tier = "DIAMOND"
)

// Division …
type Division string

//
const (
	I   Division = "I"
	II  Division = "II"
	III Division = "III"
	IV  Division = "IV"
)

const (
	leagueRootURL  = "league/v4/"
	challengerURL  = "challengerleagues/by-queue"
	grandmasterURL = "grandmasterleagues/by-queue"
	masterURL      = "masterleagues/by-queue"
	leaguesURL     = "leagues"
	entriesURL     = "entries/"
)

// LeagueEntryDTO …
type LeagueEntryDTO struct {
	SummonerName string        `json:"summonerName"`
	SummonerID   string        `json:"summonerId"`
	LeagueID     string        `json:"leagueId"`
	QueueType    string        `json:"queueType"`
	Tier         string        `json:"tier"`
	Rank         string        `json:"rank"`
	LeaguePoints int           `json:"leaguePoints"`
	Wins         int           `json:"wins"`
	Losses       int           `json:"losses"`
	HotStreak    bool          `json:"hotStreak"`
	MiniSeries   MiniSeriesDTO `json:"miniSeries"`
	FreshBlood   bool          `json:"freshBlood"`
	Veteran      bool          `json:"Veteran"`
	Inactive     bool          `json:"inactive"`
}

// LeagueListDTO …
type LeagueListDTO struct {
	LeagueID string          `json:"leagueId"`
	Queue    string          `json:"queue"`
	Name     string          `json:"name"`
	Tier     string          `json:"tier"`
	Entries  []LeagueItemDTO `json:"entries"`
}

// LeagueItemDTO …
type LeagueItemDTO struct {
	SummonerName string        `json:"summonerName"`
	SummonerID   string        `json:"summonerID"`
	Rank         string        `json:"rank"`
	LeaguePoints int           `json:"leaguePoints"`
	Wins         int           `json:"wins"`
	Losses       int           `json:"losses"`
	HotStreak    bool          `json:"hotStreak"`
	MiniSeries   MiniSeriesDTO `json:"miniSeries"`
	FreshBlood   bool          `json:"freshBlood"`
	Veteran      bool          `json:"veteran"`
	Inactive     bool          `json:"inactive"`
}

// MiniSeriesDTO …
type MiniSeriesDTO struct {
	Wins     int    `json:"wins"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Progress string `json:"progress"`
}

// ChallengerLeagueByQueue gets challenger leagues for the given queue type
func (c *client) ChallengerLeagueByQueue(ctx context.Context, queue QueueType) (LeagueListDTO, error) {
	var res LeagueListDTO
	url := filepath.Join(leagueRootURL, challengerURL, string(queue))
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// GrandmasterLeagueByQueue gets grandmaster leagues for the given queue type
func (c *client) GrandmasterLeagueByQueue(ctx context.Context, queue QueueType) (LeagueListDTO, error) {
	var res LeagueListDTO
	url := filepath.Join(leagueRootURL, grandmasterURL, string(queue))
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// MasterLeagueByQueue gets master leagues for the given queue type
func (c *client) MasterLeagueByQueue(ctx context.Context, queue QueueType) (LeagueListDTO, error) {
	var res LeagueListDTO
	url := filepath.Join(leagueRootURL, masterURL, string(queue))
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Leagues gets the leagues for the given leagueID
func (c *client) Leagues(ctx context.Context, leagueID string) (LeagueListDTO, error) {
	var res LeagueListDTO
	url := filepath.Join(leagueRootURL, leaguesURL, leagueID)
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Entries returns the entries in the given queue at the given tier and given division
func (c *client) Entries(ctx context.Context, queue QueueType, tier Tier, division Division) ([]LeagueEntryDTO, error) {
	var res []LeagueEntryDTO
	url := filepath.Join(leagueRootURL, entriesURL, string(queue), string(tier), string(division))
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// EntriesBySummoner returns the league entries for the given summoner ID
func (c *client) EntriesBySummoner(ctx context.Context, summonerID string) ([]LeagueEntryDTO, error) {
	var res []LeagueEntryDTO
	url := filepath.Join(leagueRootURL, entriesURL, "by-summoner", summonerID)
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (le LeagueEntryDTO) String() string {
	var str strings.Builder
	str.WriteString("LeagueEntryDTO{\n")
	str.WriteString(fmt.Sprintf("    SummonerName: \"%v\",\n", le.SummonerName))
	str.WriteString(fmt.Sprintf("    SummonerID: \"%v\",\n", le.SummonerID))
	str.WriteString(fmt.Sprintf("    LeagueID: \"%v\",\n", le.LeagueID))
	str.WriteString(fmt.Sprintf("    QueueType: \"%v\",\n", le.QueueType))
	str.WriteString(fmt.Sprintf("    Tier: \"%v\",\n", le.Tier))
	str.WriteString(fmt.Sprintf("    Rank: \"%v\",\n", le.Rank))
	str.WriteString(fmt.Sprintf("    LeaguePoints: %v,\n", le.LeaguePoints))
	str.WriteString(fmt.Sprintf("    Wins: %v,\n", le.Wins))
	str.WriteString(fmt.Sprintf("    Losses: %v,\n", le.Losses))
	str.WriteString(fmt.Sprintf("    HotStreak: %v,\n", le.HotStreak))
	str.WriteString(fmt.Sprintf("    MiniSeries: %v,\n", le.MiniSeries))
	str.WriteString(fmt.Sprintf("    FreshBlood: %v,\n", le.FreshBlood))
	str.WriteString(fmt.Sprintf("    Veteran: %v,\n", le.Veteran))
	str.WriteString(fmt.Sprintf("    Inactive: %v,\n", le.Inactive))
	str.WriteString("},\n")
	return str.String()
}

func (ll LeagueListDTO) String() string {
	var str strings.Builder
	str.WriteString("LeagueListDTO{\n")
	str.WriteString(fmt.Sprintf("    LeagueID: \"%v\",\n", ll.LeagueID))
	str.WriteString(fmt.Sprintf("    Queue: \"%v\",\n", ll.Queue))
	str.WriteString(fmt.Sprintf("    Name: \"%v\",\n", ll.Name))
	str.WriteString(fmt.Sprintf("    Tier: \"%v\",\n", ll.Tier))
	str.WriteString(fmt.Sprintf("    Entries: %v,\n", ll.Entries))
	str.WriteString("},\n")
	return str.String()
}

func (li LeagueItemDTO) String() string {
	var str strings.Builder
	str.WriteString("LeagueItemDTO{\n")
	str.WriteString(fmt.Sprintf("    SummonerName: \"%v\",\n", li.SummonerName))
	str.WriteString(fmt.Sprintf("    SummonerID: \"%v\",\n", li.SummonerName))
	str.WriteString(fmt.Sprintf("    Rank: \"%v\",\n", li.Rank))
	str.WriteString(fmt.Sprintf("    LeaguePoints: %v,\n", li.LeaguePoints))
	str.WriteString(fmt.Sprintf("    Wins: %v,\n", li.Wins))
	str.WriteString(fmt.Sprintf("    Losses: %v,\n", li.Losses))
	str.WriteString(fmt.Sprintf("    HotStreak: %v,\n", li.HotStreak))
	str.WriteString(fmt.Sprintf("    MiniSeries: %v,\n", li.MiniSeries))
	str.WriteString(fmt.Sprintf("    FreshBlood: %v,\n", li.FreshBlood))
	str.WriteString(fmt.Sprintf("    Veteran: %v,\n", li.Veteran))
	str.WriteString(fmt.Sprintf("    Inactive: %v,\n", li.Inactive))
	str.WriteString("},\n")
	return str.String()
}

func (ms MiniSeriesDTO) String() string {
	var str strings.Builder
	str.WriteString("MiniSeriesDTO{\n")
	str.WriteString(fmt.Sprintf("    Wins: %v,\n", ms.Wins))
	str.WriteString(fmt.Sprintf("    Losses: %v,\n", ms.Losses))
	str.WriteString(fmt.Sprintf("    Target: %v,\n", ms.Target))
	str.WriteString(fmt.Sprintf("    Progress: \"%v\",\n", ms.Progress))
	str.WriteString("},\n")
	return str.String()
}
