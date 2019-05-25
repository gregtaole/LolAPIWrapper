package lolapiwrapper

import (
	"context"
	"fmt"
)

/*CurrentGameInfoDTO …*/
type CurrentGameInfoDTO struct {
}

/*FeaturedGamesDTO …*/
type FeaturedGamesDTO struct {
}

// ActiveGamesBySummoner gets the current game information for the given summonerID
func (c *client) ActiveGamesBySummoner(ctx context.Context, summonerID string) (CurrentGameInfoDTO, error) {
	var res CurrentGameInfoDTO
	return res, fmt.Errorf("not implemented")
}

// FeaturedGames gets a list of featured games
func (c *client) FeaturedGames(ctx context.Context) (FeaturedGamesDTO, error) {
	var res FeaturedGamesDTO
	return res, fmt.Errorf("not implemented")
}
