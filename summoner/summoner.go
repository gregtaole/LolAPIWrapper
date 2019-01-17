package summoner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gregtaole/lolapiwrapper/util"
)

const (
	rootURL      = "summoner/v4/summoners/"
	byAccountURL = "by-account"
	byNameURL    = "by-name"
	byPuuidURL   = "by-puuid"
)

type Summoner struct {
	APIKey string
	Region string
}

type SummonerDTO struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	Puuid         string `json:"puuid"`
	LummonerLevel int    `json:"summonerLevel"`
	RevisionDate  int    `json:"revisionDate"`
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
}

func NewSummoner(APIKey, region string) Summoner {
	return Summoner{
		APIKey: APIKey,
		Region: region,
	}
}

func (s Summoner) SummonerByAccount(accountID string) (*SummonerDTO, error) {
	resp, err := util.GetResponse(s.APIKey, s.Region, rootURL+byAccountURL, accountID)
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
	var summonerDTO SummonerDTO
	if err = json.Unmarshal(body, &summonerDTO); err != nil {
		return nil, err
	}
	return &summonerDTO, nil
}

func (s Summoner) SummonerByName(name string) (*SummonerDTO, error) {
	resp, err := util.GetResponse(s.APIKey, s.Region, rootURL+byNameURL, name)
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
	var summonerDTO SummonerDTO
	if err = json.Unmarshal(body, &summonerDTO); err != nil {
		return nil, err
	}
	return &summonerDTO, nil
}

func (s Summoner) SummonerByPuuid(puuid string) (*SummonerDTO, error) {
	resp, err := util.GetResponse(s.APIKey, s.Region, rootURL+byPuuidURL, puuid)
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
	var summonerDTO SummonerDTO
	if err = json.Unmarshal(body, &summonerDTO); err != nil {
		return nil, err
	}
	return &summonerDTO, nil
}

func (s Summoner) SummonerByID(ID string) (*SummonerDTO, error) {
	resp, err := util.GetResponse(s.APIKey, s.Region, "summoner/v4/summoners", ID)
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
	var summonerDTO SummonerDTO
	if err = json.Unmarshal(body, &summonerDTO); err != nil {
		return nil, err
	}
	return &summonerDTO, nil
}

func (s SummonerDTO) String() string {
	return fmt.Sprintf("%v, %v", s.Name, s.Puuid)
}
