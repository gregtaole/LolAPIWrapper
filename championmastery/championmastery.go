package championmastery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gregtaole/lolapiwrapper/util"
)

const (
	rootURL       = "champion-mastery/v4/"
	bySummonerURL = "champion-masteries/by-summoner"
	byChampionURL = "by-champion"
	scoresURL     = "scores/by-summoner"
)

type ChampionMastery struct {
	APIKey string
	Region string
}

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

func NewChampionMastery(APIKey, region string) ChampionMastery {
	return ChampionMastery{
		APIKey: APIKey,
		Region: region,
	}
}

func (cm ChampionMastery) ChampionMasteriesBySummoner(summonerID string) ([]ChampionMasteryDTO, error) {
	resp, err := util.GetResponse(cm.APIKey, cm.Region, rootURL+bySummonerURL, summonerID)
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
	var championMasteryDTOs []ChampionMasteryDTO
	if err = json.Unmarshal(body, &championMasteryDTOs); err != nil {
		return nil, err
	}
	return championMasteryDTOs, nil
}

func (cm ChampionMastery) ChampionMasteriesBySummonerByChampion(summonerID string, championID int) (*ChampionMasteryDTO, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://%v.api.riotgames.com/lol/%v/%v/%v/%v", cm.Region, rootURL+bySummonerURL, summonerID, byChampionURL, championID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request for %v: %v", url, err)
	}
	fmt.Println(req.URL)
	req.Header.Set("X-Riot-Token", cm.APIKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not get %v: %v", req.URL, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server responded with error code %v", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var championMasteryDTO ChampionMasteryDTO
	if err = json.Unmarshal(body, &championMasteryDTO); err != nil {
		return nil, err
	}
	return &championMasteryDTO, nil
}

func (cm ChampionMastery) ScoresBySummoner(summonerID string) (*int, error) {
	resp, err := util.GetResponse(cm.APIKey, cm.Region, rootURL+scoresURL, summonerID)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server responded with error code %v", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var masteryScore int
	if err = json.Unmarshal(body, &masteryScore); err != nil {
		return nil, err
	}
	return &masteryScore, nil
}

func (cmd ChampionMasteryDTO) String() string {
	return fmt.Sprintf("%v\n", cmd.ChampionLevel)
}
