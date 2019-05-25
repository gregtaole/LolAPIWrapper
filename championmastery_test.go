package lolapiwrapper

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestChampionMasteriesBySummoner(t *testing.T) {
	limiter := rate.NewLimiter(20, 1)
	doer := http.DefaultClient
	c := client{
		HTTPClient: doer,
		Limiter:    limiter,
	}
	var tests = []struct {
		summonerID string
		want       []ChampionMasteryDTO
	}{
		{"6RyT_m6Bn-iUd1n3xkOthI1BVpuLjWj8aUMD8nrJogfMpUrZ", []ChampionMasteryDTO{
			ChampionMasteryDTO{
				ChestGranted:                 true,
				ChampionLevel:                5,
				ChampionPoints:               112585,
				ChampionID:                   142,
				ChampionPointsUntilNextLevel: 0,
				ChampionPointsSinceLastLevel: 90985,
				LastPlayTime:                 1543619081000,
				TokensEarned:                 0,
				SummonerID:                   "6RyT_m6Bn-iUd1n3xkOthI1BVpuLjWj8aUMD8nrJogfMpUrZ",
			},
			ChampionMasteryDTO{
				ChestGranted:                 true,
				ChampionLevel:                5,
				ChampionPoints:               32496,
				ChampionID:                   203,
				ChampionPointsUntilNextLevel: 0,
				ChampionPointsSinceLastLevel: 10896,
				LastPlayTime:                 1544888996000,
				TokensEarned:                 0,
				SummonerID:                   "6RyT_m6Bn-iUd1n3xkOthI1BVpuLjWj8aUMD8nrJogfMpUrZ",
			},
			ChampionMasteryDTO{
				ChestGranted:                 true,
				ChampionLevel:                4,
				ChampionPoints:               17878,
				ChampionID:                   429,
				ChampionPointsUntilNextLevel: 3722,
				ChampionPointsSinceLastLevel: 5278,
				LastPlayTime:                 1539789673000,
				TokensEarned:                 0,
				SummonerID:                   "6RyT_m6Bn-iUd1n3xkOthI1BVpuLjWj8aUMD8nrJogfMpUrZ",
			},
		}},
	}

	for _, test := range tests {
		jsonData, err := ioutil.ReadFile("test_data/championmastery/by_summoner/" + test.summonerID + ".json")
		if err != nil {
			t.Fatalf("could not open test data file: %v", err)
		}
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "%s", jsonData)
		}))
		defer server.Close()
		c.URLBuilder = func(string, string, string) string { return server.URL }
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		got, err := c.ChampionMasteriesBySummoner(ctx, test.summonerID)
		if err != nil {
			t.Errorf("query returned with error: %v", err)
		}
		if !ChampionMasteryDTOSliceEqual(got, test.want) {
			t.Errorf("%v, want %v", got, test.want)
		}
	}

}

func ChampionMasteryDTOSliceEqual(a, b []ChampionMasteryDTO) bool {
	if len(a) != len(b) {
		fmt.Println("lengths do not match")
		return false
	}
	for i, v := range a {
		if v != b[i] {
			fmt.Println("elements do not match")
			return false
		}
	}
	return true
}
