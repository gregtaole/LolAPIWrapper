package lolapiwrapper

import (
	"context"
	"fmt"
)

/*QueueType …*/
type QueueType string

/* …*/
const (
	Solo   QueueType = "RANKED_SOLO_5x5"
	FlexSR QueueType = "RANKED_FLEX_SR"
	FlexTT QueueType = "RANKED_FLEX_TT"
)

const (
	leagueRootURL  = "league/v4/"
	challengerURL  = "challengerleagues/by-queue"
	grandmasterURL = "grandmasterleagues/by-queue"
	leaguesURL     = "leagues"
	masterURL      = "masterleagues/by-queue"
	positionsURL   = "positions/by-summoner"
)

/*LeagueListDTO … */
type LeagueListDTO struct {
	LeagueID string          `json:"leagueId"`
	Tier     string          `json:"tier"`
	Entries  []LeagueItemDTO `json:"entries"`
	Queue    string          `json:"queue"`
	Name     string          `json:"name"`
}

/*LeagueItemDTO … */
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

/*LeaguePositionDTO … */
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

/*MiniSeriesDTO … */
type MiniSeriesDTO struct {
	Wins     int    `json:"wins"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Progress string `json:"progress"`
}

// ChallengerLeagueByQueue gets challenger leagues for the given queue type
func (c *client) ChallengerLeagueByQueue(ctx context.Context, queue QueueType) (LeagueListDTO, error) {
	/* resp, err := util.GetResponse(l.APIKey, l.Region, rootURL+challengerURL, string(queue))
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
	return &leagueListDTO, nil */
	var res LeagueListDTO
	return res, fmt.Errorf("not implemented")
}

// GrandmasterLeagueByQueue gets grandmaster leagues for the given queue type
func (c *client) GrandmasterLeagueByQueue(ctx context.Context, queue QueueType) (LeagueListDTO, error) {
	/* resp, err := util.GetResponse(l.APIKey, l.Region, rootURL+grandmasterURL, string(queue))
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
	return &leagueListDTO, nil */
	var res LeagueListDTO
	return res, fmt.Errorf("not implemented")
}

// MasterLeagueByQueue gets master leagues for the given queue type
func (c *client) MasterLeagueByQueue(ctx context.Context, queue QueueType) (LeagueListDTO, error) {
	/* resp, err := util.GetResponse(l.APIKey, l.Region, rootURL+masterURL, string(queue))
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
	return &leagueListDTO, nil */
	var res LeagueListDTO
	return res, fmt.Errorf("not implemented")
}

// PositionsBySummoner gets the positions for summonerID
func (c *client) PositionsBySummoner(ctx context.Context, summonerID string) ([]LeaguePositionDTO, error) {
	/* resp, err := util.GetResponse(l.APIKey, l.Region, rootURL+positionsURL, summonerID)
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
	return leaguePositionDTO, nil */
	return nil, fmt.Errorf("not implemented")
}

// Leagues gets the leagues for the given leagueID
func (c *client) Leagues(ctx context.Context, leagueID string) (LeagueListDTO, error) {
	/* resp, err := util.GetResponse(l.APIKey, l.Region, rootURL+leaguesURL, leagueID)
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
	return &leagueListDTO, nil */
	var res LeagueListDTO
	return res, fmt.Errorf("not implemented")
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
