package lolapiwrapper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

/*
Doer is the interface to execute http requests
*/
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

/*
Limiter is an interface for the rate limiter
*/
type Limiter interface {
	Wait(context.Context) error
}

/*
URLBuilder is an interface to build the full URL by providing the region, the query url and the parameters
*/
type URLBuilder func(string, string, string) string

/*
Client is the interface to access the Riot API
*/
type Client interface {
	// Champion Mastery API

	// ChampionMasteriesBySummoner gets all champion masteries for summonerID
	ChampionMasteriesBySummoner(ctx context.Context, summonerID string) ([]ChampionMasteryDTO, error)
	// ChampionMasteriesBySummonerByChampion gets the mastery score for championID for summonerID
	ChampionMasteriesBySummonerByChampion(ctx context.Context, summonerID string, championID int) (ChampionMasteryDTO, error)
	// ScoresBySummoner gets all mastery scores for summonerID
	ScoresBySummoner(ctx context.Context, summonerID string) (int, error)

	// League API

	// ChallengerLeagueByQueue gets challenger leagues for the given queue type
	ChallengerLeagueByQueue(ctx context.Context, queue QueueType) (LeagueListDTO, error)
	// GrandmasterLeagueByQueue gets grandmaster leagues for the given queue type
	GrandmasterLeagueByQueue(ctx context.Context, queue QueueType) (LeagueListDTO, error)
	// MasterLeagueByQueue gets master leagues for the given queue type
	MasterLeagueByQueue(ctx context.Context, queue QueueType) (LeagueListDTO, error)
	// Leagues gets the leagues for the given leagueID
	Leagues(ctx context.Context, leagueID string) (LeagueListDTO, error)
	//Entries returns the entries in the given queue at the given tier and given division
	Entries(ctx context.Context, queue QueueType, tier Tier, division Division) ([]LeagueEntryDTO, error)
	//EntriesBySummoner returns the league entries for the given summoner ID
	EntriesBySummoner(ctx context.Context, summonerID string) ([]LeagueEntryDTO, error)

	// Match API

	// MatchesByID gets the match information for the given matchID
	MatchesByID(ctx context.Context, matchID string) (MatchDTO, error)
	// MatchListByAccount gets the match list for the given accountID filtered by params
	MatchListByAccount(ctx context.Context, accountID string, params MatchQueryParams) (MatchListDTO, error)
	//TimelineByMatch gets the match timeline for the given matchID
	TimelineByMatch(ctx context.Context, matchID string) (MatchTimelineDTO, error)

	// Summoner API

	// SummonerByAccount gets summoner information for the given accountID
	SummonerByAccount(ctx context.Context, accountID string) (SummonerDTO, error)
	// SummonerByName gets summoner information for the given name
	SummonerByName(ctx context.Context, name string) (SummonerDTO, error)
	// SummonerByPuuid gets summoner information for the given puuid
	SummonerByPuuid(ctx context.Context, puuid string) (SummonerDTO, error)
	// SummonerByID gets summoner informatino for the given ID
	SummonerByID(ctx context.Context, ID string) (SummonerDTO, error)

	// Spectator API

	// ActiveGamesBySummoner gets the current game information for the given summonerID
	ActiveGamesBySummoner(ctx context.Context, summonerID string) (CurrentGameInfoDTO, error)
	// FeaturedGames gets a list of featured games
	FeaturedGames(ctx context.Context) (FeaturedGamesDTO, error)
}

type client struct {
	APIKey     string
	Region     string
	HTTPClient Doer
	Limiter    Limiter
	URLBuilder URLBuilder
}

/*
MatchQueryParams is a struct containing the query parameters for the MatchListByAccount method.
*/
type MatchQueryParams struct {
	Champion   []int
	Queue      []int
	Season     []int
	EndTime    *int
	BeginTime  *int
	EndIndex   *int
	BeginIndex *int
}

/*
NewClient returns a Client to access the Riot API
*/
func NewClient(APIKey string, region string, httpClient Doer, limiter Limiter) Client {
	return &client{
		APIKey:     APIKey,
		Region:     region,
		HTTPClient: httpClient,
		Limiter:    limiter,
		URLBuilder: buildURL,
	}
}

func (c *client) query(ctx context.Context, url string, params url.Values, res interface{}) error {
	var fullURL, suffix string
	if params != nil {
		suffix = fmt.Sprintf("?%s", params.Encode())
	}
	fullURL = c.URLBuilder(c.Region, url, suffix)
	//fmt.Println(fullURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Riot-Token", c.APIKey)
	req = req.WithContext(ctx)

	err = c.Limiter.Wait(ctx)
	if err != nil {
		return err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server responded with error code %v", resp.StatusCode)
	}
	/* body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, res) */
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&res)

	return err
}

func buildURL(region, url, suffix string) string {
	return fmt.Sprintf("https://%v.api.riotgames.com/lol/%v%v", region, url, suffix)
}
