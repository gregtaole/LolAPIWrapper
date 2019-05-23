package lolapiwrapper

import (
	"context"
	"fmt"
)

type CurrentGameInfoDTO struct {
}

type FeaturedGamesDTO struct {
}

// ActiveGamesBySummoner gets the current game information for the given summonerID
func (c *client) ActiveGamesBySummoner(ctx context.Context, summonerID string) (*CurrentGameInfoDTO, error) {
	return nil, fmt.Errorf("not implemented")
}

// FeaturedGames gets a list of featured games
func (c *client) FeaturedGames(ctx context.Context) (*FeaturedGamesDTO, error) {
	return nil, fmt.Errorf("not implemented")
}
