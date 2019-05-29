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
				ChampionPoints:               112756,
				ChampionID:                   142,
				ChampionPointsUntilNextLevel: 0,
				ChampionPointsSinceLastLevel: 91156,
				LastPlayTime:                 1543619081000,
				TokensEarned:                 0,
				SummonerID:                   "6RyT_m6Bn-iUd1n3xkOthI1BVpuLjWj8aUMD8nrJogfMpUrZ",
			},
			ChampionMasteryDTO{
				ChestGranted:                 true,
				ChampionLevel:                5,
				ChampionPoints:               32637,
				ChampionID:                   203,
				ChampionPointsUntilNextLevel: 0,
				ChampionPointsSinceLastLevel: 11037,
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
		{"_JEwwuOSAIZuxsZYp1V3mObKjIWjle8eAZgXWfoksGns74U", []ChampionMasteryDTO{
			ChampionMasteryDTO{
				ChestGranted:                 true,
				ChampionLevel:                7,
				ChampionPoints:               381669,
				ChampionID:                   161,
				ChampionPointsUntilNextLevel: 0,
				ChampionPointsSinceLastLevel: 360069,
				LastPlayTime:                 1537114306000,
				TokensEarned:                 0,
				SummonerID:                   "_JEwwuOSAIZuxsZYp1V3mObKjIWjle8eAZgXWfoksGns74U",
			},
			ChampionMasteryDTO{
				ChestGranted:                 true,
				ChampionLevel:                7,
				ChampionPoints:               305119,
				ChampionID:                   111,
				ChampionPointsUntilNextLevel: 0,
				ChampionPointsSinceLastLevel: 283519,
				LastPlayTime:                 1537115899000,
				TokensEarned:                 0,
				SummonerID:                   "_JEwwuOSAIZuxsZYp1V3mObKjIWjle8eAZgXWfoksGns74U",
			},
			ChampionMasteryDTO{
				ChestGranted:                 true,
				ChampionLevel:                6,
				ChampionPoints:               60881,
				ChampionID:                   53,
				ChampionPointsUntilNextLevel: 0,
				ChampionPointsSinceLastLevel: 39281,
				LastPlayTime:                 1535393690000,
				TokensEarned:                 0,
				SummonerID:                   "_JEwwuOSAIZuxsZYp1V3mObKjIWjle8eAZgXWfoksGns74U",
			},
		}},
		{"eolRXteXeqUQpy6lVgGSGYMGs7ht9cWFkvAlNUtOQ_01UzQ", []ChampionMasteryDTO{
			ChampionMasteryDTO{
				ChestGranted:                 false,
				ChampionLevel:                3,
				ChampionPoints:               10693,
				ChampionID:                   76,
				ChampionPointsUntilNextLevel: 1907,
				ChampionPointsSinceLastLevel: 4693,
				LastPlayTime:                 1526849163000,
				TokensEarned:                 0,
				SummonerID:                   "eolRXteXeqUQpy6lVgGSGYMGs7ht9cWFkvAlNUtOQ_01UzQ",
			},
			ChampionMasteryDTO{
				ChestGranted:                 false,
				ChampionLevel:                3,
				ChampionPoints:               9896,
				ChampionID:                   22,
				ChampionPointsUntilNextLevel: 2704,
				ChampionPointsSinceLastLevel: 3896,
				LastPlayTime:                 1530821017000,
				TokensEarned:                 0,
				SummonerID:                   "eolRXteXeqUQpy6lVgGSGYMGs7ht9cWFkvAlNUtOQ_01UzQ",
			},
			ChampionMasteryDTO{
				ChestGranted:                 false,
				ChampionLevel:                3,
				ChampionPoints:               7559,
				ChampionID:                   245,
				ChampionPointsUntilNextLevel: 5041,
				ChampionPointsSinceLastLevel: 1559,
				LastPlayTime:                 1534457220000,
				TokensEarned:                 0,
				SummonerID:                   "eolRXteXeqUQpy6lVgGSGYMGs7ht9cWFkvAlNUtOQ_01UzQ",
			},
		}},
		{"hv02SPkEwbGjIpwBhULUWj3ksoEkdTq3R7SSdmV-BKGrjDU", []ChampionMasteryDTO{
			ChampionMasteryDTO{
				ChestGranted:                 true,
				ChampionLevel:                7,
				ChampionPoints:               327836,
				ChampionID:                   67,
				ChampionPointsUntilNextLevel: 0,
				ChampionPointsSinceLastLevel: 306236,
				LastPlayTime:                 1537292391000,
				TokensEarned:                 0,
				SummonerID:                   "hv02SPkEwbGjIpwBhULUWj3ksoEkdTq3R7SSdmV-BKGrjDU",
			},
			ChampionMasteryDTO{
				ChestGranted:                 false,
				ChampionLevel:                7,
				ChampionPoints:               165050,
				ChampionID:                   245,
				ChampionPointsUntilNextLevel: 0,
				ChampionPointsSinceLastLevel: 143450,
				LastPlayTime:                 1532097472000,
				TokensEarned:                 0,
				SummonerID:                   "hv02SPkEwbGjIpwBhULUWj3ksoEkdTq3R7SSdmV-BKGrjDU",
			},
			ChampionMasteryDTO{
				ChestGranted:                 false,
				ChampionLevel:                7,
				ChampionPoints:               150177,
				ChampionID:                   202,
				ChampionPointsUntilNextLevel: 0,
				ChampionPointsSinceLastLevel: 128577,
				LastPlayTime:                 1535564280000,
				TokensEarned:                 0,
				SummonerID:                   "hv02SPkEwbGjIpwBhULUWj3ksoEkdTq3R7SSdmV-BKGrjDU",
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
			t.Fatalf("query returned with error: %v", err)
		}
		if !ChampionMasteryDTOSliceEqual(got, test.want) {
			t.Errorf("%v, want %v", got, test.want)
		}
	}
}

func TestChampionMasteriesBySummonerByChampion(t *testing.T) {
	limiter := rate.NewLimiter(20, 1)
	doer := http.DefaultClient
	c := client{
		HTTPClient: doer,
		Limiter:    limiter,
	}
	var tests = []struct {
		summonerID string
		championID int
		want       ChampionMasteryDTO
	}{
		{"6RyT_m6Bn-iUd1n3xkOthI1BVpuLjWj8aUMD8nrJogfMpUrZ", 142, ChampionMasteryDTO{
			ChampionID:                   142,
			ChampionLevel:                5,
			ChampionPoints:               112756,
			ChampionPointsSinceLastLevel: 91156,
			ChampionPointsUntilNextLevel: 0,
			ChestGranted:                 true,
			LastPlayTime:                 1543619081000,
			TokensEarned:                 0,
			SummonerID:                   "6RyT_m6Bn-iUd1n3xkOthI1BVpuLjWj8aUMD8nrJogfMpUrZ",
		}},
		{"_JEwwuOSAIZuxsZYp1V3mObKjIWjle8eAZgXWfoksGns74U", 142, ChampionMasteryDTO{
			ChampionID:                   142,
			ChampionLevel:                1,
			ChampionPoints:               363,
			ChampionPointsSinceLastLevel: 363,
			ChampionPointsUntilNextLevel: 1437,
			ChestGranted:                 false,
			LastPlayTime:                 1531518060000,
			TokensEarned:                 0,
			SummonerID:                   "_JEwwuOSAIZuxsZYp1V3mObKjIWjle8eAZgXWfoksGns74U",
		}},
		{"hv02SPkEwbGjIpwBhULUWj3ksoEkdTq3R7SSdmV-BKGrjDU", 142, ChampionMasteryDTO{
			ChampionID:                   142,
			ChampionLevel:                7,
			ChampionPoints:               47025,
			ChampionPointsSinceLastLevel: 25425,
			ChampionPointsUntilNextLevel: 0,
			ChestGranted:                 true,
			LastPlayTime:                 1533480452000,
			TokensEarned:                 0,
			SummonerID:                   "hv02SPkEwbGjIpwBhULUWj3ksoEkdTq3R7SSdmV-BKGrjDU",
		}},
		{"eolRXteXeqUQpy6lVgGSGYMGs7ht9cWFkvAlNUtOQ_01UzQ", 142, ChampionMasteryDTO{
			ChampionID:                   142,
			ChampionLevel:                2,
			ChampionPoints:               2726,
			ChampionPointsSinceLastLevel: 926,
			ChampionPointsUntilNextLevel: 3274,
			ChestGranted:                 false,
			LastPlayTime:                 1535319687000,
			TokensEarned:                 0,
			SummonerID:                   "eolRXteXeqUQpy6lVgGSGYMGs7ht9cWFkvAlNUtOQ_01UzQ",
		}},
	}

	for _, test := range tests {
		jsonData, err := ioutil.ReadFile("test_data/championmastery/by_champion/" + test.summonerID + ".json")
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
		got, err := c.ChampionMasteriesBySummonerByChampion(ctx, test.summonerID, test.championID)
		if err != nil {
			t.Fatalf("query returned with error: %v", err)
		}
		if got != test.want {
			t.Errorf("%v, want %v", got, test.want)
		}
	}
}

func TestScoresBySummoner(t *testing.T) {
	limiter := rate.NewLimiter(20, 1)
	doer := http.DefaultClient
	c := client{
		HTTPClient: doer,
		Limiter:    limiter,
	}
	var tests = []struct {
		summonerID string
		want       int
	}{
		{"_JEwwuOSAIZuxsZYp1V3mObKjIWjle8eAZgXWfoksGns74U", 182},
		{"6RyT_m6Bn-iUd1n3xkOthI1BVpuLjWj8aUMD8nrJogfMpUrZ", 43},
		{"hv02SPkEwbGjIpwBhULUWj3ksoEkdTq3R7SSdmV-BKGrjDU", 418},
		{"eolRXteXeqUQpy6lVgGSGYMGs7ht9cWFkvAlNUtOQ_01UzQ", 80},
	}
	for _, test := range tests {
		jsonData, err := ioutil.ReadFile("test_data/championmastery/scores/" + test.summonerID + ".json")
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
		got, err := c.ScoresBySummoner(ctx, test.summonerID)
		if err != nil {
			t.Fatalf("query returned with error: %v", err)
		}
		if got != test.want {
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
