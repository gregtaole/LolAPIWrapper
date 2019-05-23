package lolapiwrapper

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

var apiKey = os.Getenv("API_KEY")
var limiter = rate.NewLimiter(20, 1)
var doer = http.DefaultClient
var c = NewClient(apiKey, "euw1", doer, limiter)

func TestChampionMasteriesBySummoner(t *testing.T) {
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
	type Data struct {
		ID   string
		data string
	}
	file, err := os.Open("championmastery_test.json")
	if err != nil {
		t.Fatalf("could not find test data: %v", err)
	}
	defer file.Close()
	dataMap := make(map[string]Data)
	decoder := json.NewDecoder(bufio.NewReader(file))
	for {
		var datum Data
		if err := decoder.Decode(&datum); err == io.EOF {
			break
		} else if err != nil {
			t.Fatalf("could not decode test data: %v", err)
		}
		dataMap[datum.ID] = datum
	}
	for _, test := range tests {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, dataMap[test.summonerID].data)
		}))
		defer server.Close()
		cm := ChampionMasteryDTO{}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		got, err := c.ChampionMasteriesBySummoner(ctx, test.summonerID)
		if err != nil {
			t.Errorf("query returned with error: %v", err)
		}
		if ChampionMasteryDTOSliceEqual(*got, test.want) {
			t.Errorf("%v = %v, want %v", cm, got, test.want)
		}
	}

}

func ChampionMasteryDTOSliceEqual(a, b []ChampionMasteryDTO) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
