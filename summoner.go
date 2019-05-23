package lolapiwrapper

import (
	"context"
	"fmt"
	"path/filepath"
)

const (
	summonerRootURL = "summoner/v4/summoners/"
	byAccountURL    = "by-account"
	byNameURL       = "by-name"
	byPuuidURL      = "by-puuid"
)

type SummonerDTO struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	Puuid         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
	RevisionDate  int    `json:"revisionDate"`
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
}

// SummonerByAccount gets summoner information for the given accountID
func (c *client) SummonerByAccount(ctx context.Context, accountID string) (*SummonerDTO, error) {
	var res SummonerDTO
	url := filepath.Join(summonerRootURL, byAccountURL, accountID)
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// SummonerByName gets summoner information for the given name
func (c *client) SummonerByName(ctx context.Context, name string) (*SummonerDTO, error) {
	var res SummonerDTO
	url := filepath.Join(summonerRootURL, byNameURL, name)
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// SummonerByPuuid gets summoner information for the given puuid
func (c *client) SummonerByPuuid(ctx context.Context, puuid string) (*SummonerDTO, error) {
	var res SummonerDTO
	url := filepath.Join(summonerRootURL, byPuuidURL, puuid)
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// SummonerByID gets summoner informatino for the given ID
func (c *client) SummonerByID(ctx context.Context, ID string) (*SummonerDTO, error) {
	var res SummonerDTO
	url := filepath.Join(summonerRootURL, ID)
	err := c.query(ctx, url, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s SummonerDTO) String() string {
	return fmt.Sprintf("Name: %v,\nAccountID: %v,\nID: %v,\nPUUID: %v,\nSummonerLevel: %v,\nRevisionDate: %v,\nProfileIconID: %v", s.Name, s.AccountID, s.ID, s.Puuid, s.SummonerLevel, s.RevisionDate, s.ProfileIconID)
}
