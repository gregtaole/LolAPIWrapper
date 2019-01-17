package lolapiwrapper

import (
	"github.com/gregtaole/lolapiwrapper/championmastery"
	"github.com/gregtaole/lolapiwrapper/league"
	"github.com/gregtaole/lolapiwrapper/match"
	"github.com/gregtaole/lolapiwrapper/summoner"
)

/*
Client is the basic struct to access the API
*/
type Client struct {
	APIKey             string
	Region             string
	ChampionMasteryAPI championmastery.ChampionMastery
	LeagueAPI          league.League
	MatchAPI           match.Match
	SummonerAPI        summoner.Summoner
}

func NewClient(APIKey string, region string) Client {
	cm := championmastery.NewChampionMastery(APIKey, region)
	l := league.NewLeague(APIKey, region)
	m := match.NewMatch(APIKey, region)
	s := summoner.NewSummoner(APIKey, region)

	return Client{
		APIKey:             APIKey,
		Region:             region,
		ChampionMasteryAPI: cm,
		LeagueAPI:          l,
		MatchAPI:           m,
		SummonerAPI:        s,
	}
}
