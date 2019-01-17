package league

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gregtaole/lolapiwrapper/util"
)

type QueueType string

const (
	Solo   QueueType = "RANKED_SOLO_5x5"
	FlexSR QueueType = "RANKED_FLEX_SR"
	FlexTT QueueType = "RANKED_FLEX_TT"
)

const (
	rootURL        = "league/v4/"
	challengerURL  = "challengerleagues/by-queue"
	grandmasterURL = "grandmasterleagues/by-queue"
	leaguesURL     = "leagues"
	masterURL      = "masterleagues/by-queue"
	positionsURL   = "positions/by-summoner"
)

type League struct {
	APIKey string
	Region string
}

type LeagueListDTO struct {
	LeagueID string          `json:"leagueId"`
	Tier     string          `json:"tier"`
	Entries  []LeagueItemDTO `json:"entries"`
	Queue    string          `json:"queue"`
	Name     string          `json:"name"`
}

type LeagueItemDTO struct {
	Rank             string        `json:"rank"`
	HotStreak        bool          `json:"hotStreak"`
	MiniSeries       MiniSeriesDTO `json:"miniSeries"`
	Wins             int           `json:"wins"`
	Veteran          bool          `json:"veteran"`
	Losses           int           `json:"losses"`
	FreshBlood       bool          `json:"freshBlood"`
	PlayerOrTeamName string        `json:"playerOrTeamName"`
	Inactive         bool          `json:"inactive"`
	PlayerOrTeamID   string        `json:"playerOrTeamID"`
	LeaguePoints     int           `json:"leaguePoints"`
}

type LeaguePositionDTO struct {
	QueueType    string        `json:"queueType"`
	SummonerName string        `json:"summonerName"`
	HotStreak    bool          `json:"hotStreak"`
	MiniSeries   MiniSeriesDTO `json:"miniSeries"`
	Wins         int           `json:"wins"`
	Veteran      bool          `json:"veteran"`
	Losses       int           `json:"losses"`
	FreshBlood   bool          `json:"freshBlood"`
	LeagueID     string        `json:"leagueID"`
	Inactive     bool          `json:"inactive"`
	Rank         string        `json:"rank"`
	LeagueName   string        `json:"leagueName"`
	Tier         string        `json:"tier"`
	SummonerID   string        `json:"summonerId"`
	LeaguePoints int           `json:"leaguePoints"`
}

type MiniSeriesDTO struct {
	Wins     int    `json:"wins"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Progress string `json:"progress"`
}

func NewLeague(APIKey, region string) League {
	return League{
		APIKey: APIKey,
		Region: region,
	}
}

func (l League) ChallengerLeagueByQueue(queue QueueType) (*LeagueListDTO, error) {
	resp, err := util.GetResponse(l.APIKey, l.Region, rootURL+challengerURL, string(queue))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server responded with error code %v", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var leagueListDTO LeagueListDTO
	if err = json.Unmarshal(body, &leagueListDTO); err != nil {
		return nil, err
	}
	return &leagueListDTO, nil
}

func (l League) GrandmasterLeagueByQueue(queue QueueType) (*LeagueListDTO, error) {
	resp, err := util.GetResponse(l.APIKey, l.Region, rootURL+grandmasterURL, string(queue))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server responded with error code %v", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var leagueListDTO LeagueListDTO
	if err = json.Unmarshal(body, &leagueListDTO); err != nil {
		return nil, err
	}
	return &leagueListDTO, nil
}

func (l League) MasterLeagueByQueue(queue QueueType) (*LeagueListDTO, error) {
	resp, err := util.GetResponse(l.APIKey, l.Region, rootURL+masterURL, string(queue))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server responded with error code %v", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var leagueListDTO LeagueListDTO
	if err = json.Unmarshal(body, &leagueListDTO); err != nil {
		return nil, err
	}
	return &leagueListDTO, nil
}

func (l League) PositionsBySummoner(summonerID string) ([]LeaguePositionDTO, error) {
	resp, err := util.GetResponse(l.APIKey, l.Region, rootURL+positionsURL, summonerID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server responded with error code %v", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var leaguePositionDTO []LeaguePositionDTO
	if err = json.Unmarshal(body, &leaguePositionDTO); err != nil {
		return nil, err
	}
	return leaguePositionDTO, nil
}

func (l League) Leagues(leagueID string) (*LeagueListDTO, error) {
	resp, err := util.GetResponse(l.APIKey, l.Region, rootURL+leaguesURL, leagueID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server responded with error code %v", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var leagueListDTO LeagueListDTO
	if err = json.Unmarshal(body, &leagueListDTO); err != nil {
		return nil, err
	}
	return &leagueListDTO, nil
}

func (ll LeagueListDTO) String() string {
	return fmt.Sprintf("%v", ll.Tier)
}

func (li LeagueItemDTO) String() string {
	return fmt.Sprintf("%v", li.Rank)
}

func (ms MiniSeriesDTO) String() string {
	return fmt.Sprintf("%v", ms.Wins)
}
