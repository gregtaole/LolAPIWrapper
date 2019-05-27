package lolapiwrapper

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
)

const (
	summonerRootURL = "summoner/v4/summoners/"
	byAccountURL    = "by-account"
	byNameURL       = "by-name"
	byPuuidURL      = "by-puuid"
)

/*SummonerDTO …*/
type SummonerDTO struct {
	Name          string `json:"name"`
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int    `json:"revisionDate"`
}

// SummonerByAccount gets summoner information for the given accountID
func (c *client) SummonerByAccount(ctx context.Context, accountID string) (SummonerDTO, error) {
	var res SummonerDTO
	url := filepath.Join(summonerRootURL, byAccountURL, accountID)
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// SummonerByName gets summoner information for the given name
func (c *client) SummonerByName(ctx context.Context, name string) (SummonerDTO, error) {
	var res SummonerDTO
	url := filepath.Join(summonerRootURL, byNameURL, name)
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// SummonerByPuuid gets summoner information for the given puuid
func (c *client) SummonerByPuuid(ctx context.Context, puuid string) (SummonerDTO, error) {
	var res SummonerDTO
	url := filepath.Join(summonerRootURL, byPuuidURL, puuid)
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// SummonerByID gets summoner informatino for the given ID
func (c *client) SummonerByID(ctx context.Context, ID string) (SummonerDTO, error) {
	var res SummonerDTO
	url := filepath.Join(summonerRootURL, ID)
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s SummonerDTO) String() string {
	var str strings.Builder
	str.WriteString("SummonerDTO{\n")
	str.WriteString(fmt.Sprintf("    Name: %v,\n", s.Name))
	str.WriteString(fmt.Sprintf("    ID: %v,\n", s.ID))
	str.WriteString(fmt.Sprintf("    AccountID: %v,\n", s.AccountID))
	str.WriteString(fmt.Sprintf("    Puuid: %v,\n", s.Puuid))
	str.WriteString(fmt.Sprintf("    SummonerLevel: %v,\n", s.SummonerLevel))
	str.WriteString(fmt.Sprintf("    ProfileIconID: %v,\n", s.ProfileIconID))
	str.WriteString(fmt.Sprintf("    RevisionDate: %v,\n", s.RevisionDate))
	str.WriteString("}")
	return str.String()
}
