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

func TestChallengerLeagueByQueue(t *testing.T) {
	limiter := rate.NewLimiter(20, 1)
	doer := http.DefaultClient
	c := client{
		HTTPClient: doer,
		Limiter:    limiter,
	}
	var tests = []struct {
		queue QueueType
		want  LeagueListDTO
	}{
		{Solo, LeagueListDTO{
			LeagueID: "65ebcd4f-368c-30f6-a635-976beb0e1a8c",
			Queue:    "RANKED_SOLO_5x5",
			Name:     "Varus's Outriders",
			Tier:     "CHALLENGER",
			Entries: []LeagueItemDTO{
				LeagueItemDTO{
					SummonerName: "Zumo de Manzana",
					SummonerID:   "Zumo de Manzana",
					Rank:         "I",
					LeaguePoints: 666,
					Wins:         250,
					Losses:       205,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: false,
					Veteran:    false,
					Inactive:   false,
				},
				LeagueItemDTO{
					SummonerName: "GGA Supa",
					SummonerID:   "GGA Supa",
					Rank:         "I",
					LeaguePoints: 564,
					Wins:         348,
					Losses:       299,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: false,
					Veteran:    false,
					Inactive:   false,
				},
				LeagueItemDTO{
					SummonerName: "Kobbe",
					SummonerID:   "Kobbe",
					Rank:         "I",
					LeaguePoints: 679,
					Wins:         344,
					Losses:       293,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: false,
					Veteran:    true,
					Inactive:   false,
				},
			},
		}},
	}
	for _, test := range tests {
		jsonData, err := ioutil.ReadFile("test_data/league/challenger/" + test.want.LeagueID + ".json")
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
		got, err := c.ChallengerLeagueByQueue(ctx, test.queue)
		if err != nil {
			t.Errorf("query returned with error: %v", err)
		}
		if LeagueListEqual(got, test.want) {
			t.Errorf("%v, want %v", got, test.want)
		}
	}
}

func TestGrandmasterLeagueByQueue(t *testing.T) {
	limiter := rate.NewLimiter(20, 1)
	doer := http.DefaultClient
	c := client{
		HTTPClient: doer,
		Limiter:    limiter,
	}
	var tests = []struct {
		queue QueueType
		want  LeagueListDTO
	}{
		{Solo, LeagueListDTO{
			LeagueID: "5a5e7663-d053-37c2-ad35-e274b839050c",
			Queue:    "RANKED_SOLO_5x5",
			Name:     "Syndra's Masterminds",
			Tier:     "GRANDMASTER",
			Entries: []LeagueItemDTO{
				LeagueItemDTO{
					SummonerName: "LUNARY CHAP",
					SummonerID:   "LUNARY CHAP",
					Rank:         "I",
					LeaguePoints: 511,
					Wins:         187,
					Losses:       150,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: true,
					Veteran:    false,
					Inactive:   false,
				},
				LeagueItemDTO{
					SummonerName: "Rolinse",
					SummonerID:   "Rolinse",
					Rank:         "I",
					LeaguePoints: 275,
					Wins:         287,
					Losses:       252,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: false,
					Veteran:    false,
					Inactive:   false,
				},
				LeagueItemDTO{
					SummonerName: "Czajëkk",
					SummonerID:   "Czajëkk",
					Rank:         "I",
					LeaguePoints: 421,
					Wins:         158,
					Losses:       127,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: false,
					Veteran:    false,
					Inactive:   false,
				},
			},
		}},
	}
	for _, test := range tests {
		jsonData, err := ioutil.ReadFile("test_data/league/grandmaster/" + test.want.LeagueID + ".json")
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
		got, err := c.GrandmasterLeagueByQueue(ctx, test.queue)
		if err != nil {
			t.Fatalf("query returned with error: %v", err)
		}
		if LeagueListEqual(got, test.want) {
			t.Errorf("%v, want %v", got, test.want)
		}
	}
}

func TestMasterLeagueByQueue(t *testing.T) {
	limiter := rate.NewLimiter(20, 1)
	doer := http.DefaultClient
	c := client{
		HTTPClient: doer,
		Limiter:    limiter,
	}
	var tests = []struct {
		queue QueueType
		want  LeagueListDTO
	}{
		{Solo, LeagueListDTO{
			LeagueID: "0fb1590b-1aa7-3e2e-aeb4-dedeec98aecf",
			Queue:    "RANKED_SOLO_5x5",
			Name:     "Taric's Elementalists",
			Tier:     "MASTER",
			Entries: []LeagueItemDTO{
				LeagueItemDTO{
					SummonerName: "ChenDuxiu NB",
					SummonerID:   "ChenDuxiu NB",
					Rank:         "I",
					LeaguePoints: 76,
					Wins:         103,
					Losses:       79,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: true,
					Veteran:    false,
					Inactive:   false,
				},
				LeagueItemDTO{
					SummonerName: "Héphaïstoz",
					SummonerID:   "Héphaïstoz",
					Rank:         "I",
					LeaguePoints: 0,
					Wins:         93,
					Losses:       74,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: false,
					Veteran:    false,
					Inactive:   false,
				},
				LeagueItemDTO{
					SummonerName: "OP Tarzan",
					SummonerID:   "OP Tarzan",
					Rank:         "I",
					LeaguePoints: 21,
					Wins:         386,
					Losses:       368,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: false,
					Veteran:    false,
					Inactive:   false,
				},
			},
		}},
	}
	for _, test := range tests {
		jsonData, err := ioutil.ReadFile("test_data/league/master/" + test.want.LeagueID + ".json")
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
		got, err := c.MasterLeagueByQueue(ctx, test.queue)
		if err != nil {
			t.Fatalf("query returned with error: %v", err)
		}
		if LeagueListEqual(got, test.want) {
			t.Errorf("%v, want %v", got, test.want)
		}
	}
}

func TestLeagues(t *testing.T) {
	limiter := rate.NewLimiter(20, 1)
	doer := http.DefaultClient
	c := client{
		HTTPClient: doer,
		Limiter:    limiter,
	}
	var tests = []struct {
		leagueID string
		want     LeagueListDTO
	}{
		{"00f02fa0-4799-11e9-91df-c81f66dacb22", LeagueListDTO{
			LeagueID: "00f02fa0-4799-11e9-91df-c81f66dacb22",
			Queue:    "RANKED_SOLO_5x5",
			Name:     "Karma's Privateers",
			Tier:     "DIAMOND",
			Entries: []LeagueItemDTO{
				LeagueItemDTO{
					SummonerName: "Facebook",
					SummonerID:   "Facebook",
					Rank:         "IV",
					LeaguePoints: 17,
					Wins:         88,
					Losses:       63,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: false,
					Veteran:    false,
					Inactive:   false,
				},
				LeagueItemDTO{
					SummonerName: "Minipoppin",
					SummonerID:   "Minipoppin",
					Rank:         "IV",
					LeaguePoints: 0,
					Wins:         86,
					Losses:       82,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: false,
					Veteran:    false,
					Inactive:   false,
				},
				LeagueItemDTO{
					SummonerName: "ToxIgwaN",
					SummonerID:   "ToxIgwaN",
					Rank:         "IV",
					LeaguePoints: 2,
					Wins:         44,
					Losses:       46,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: false,
					Veteran:    false,
					Inactive:   false,
				},
			},
		}},
		{"003b8e80-2de3-11e9-84d4-c81f66dacb22", LeagueListDTO{
			LeagueID: "003b8e80-2de3-11e9-84d4-c81f66dacb22",
			Queue:    "RANKED_SOLO_5x5",
			Name:     "Tryndamere's Rogues",
			Tier:     "IRON",
			Entries: []LeagueItemDTO{
				LeagueItemDTO{
					SummonerName: "ZARDOWYTH",
					SummonerID:   "ZARDOWYTH",
					Rank:         "II",
					LeaguePoints: 66,
					Wins:         75,
					Losses:       117,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: false,
					Veteran:    false,
					Inactive:   false,
				},
				LeagueItemDTO{
					SummonerName: "LVDP",
					SummonerID:   "LVDP",
					Rank:         "II",
					LeaguePoints: 0,
					Wins:         20,
					Losses:       44,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: false,
					Veteran:    false,
					Inactive:   false,
				},
				LeagueItemDTO{
					SummonerName: "James007lp",
					SummonerID:   "James007lp",
					Rank:         "I",
					LeaguePoints: 17,
					Wins:         56,
					Losses:       67,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: false,
					Veteran:    true,
					Inactive:   false,
				},
			},
		}},
		{"006c10c0-7044-11e9-b5d3-c81f66dd0e0d", LeagueListDTO{
			LeagueID: "006c10c0-7044-11e9-b5d3-c81f66dd0e0d",
			Queue:    "RANKED_SOLO_5x5",
			Name:     "Irelia's Shadows",
			Tier:     "GOLD",
			Entries: []LeagueItemDTO{
				LeagueItemDTO{
					SummonerName: "RitoGomesz",
					SummonerID:   "RitoGomesz",
					Rank:         "IV",
					LeaguePoints: 21,
					Wins:         154,
					Losses:       133,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: false,
					Veteran:    false,
					Inactive:   false,
				},
				LeagueItemDTO{
					SummonerName: "Misheeru",
					SummonerID:   "Misheeru",
					Rank:         "IV",
					LeaguePoints: 0,
					Wins:         54,
					Losses:       53,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   0,
						Target:   0,
						Progress: "",
					},
					FreshBlood: false,
					Veteran:    false,
					Inactive:   false,
				},
				LeagueItemDTO{
					SummonerName: "Kollektivet",
					SummonerID:   "Kollektivet",
					Rank:         "II",
					LeaguePoints: 100,
					Wins:         5,
					Losses:       7,
					HotStreak:    false,
					MiniSeries: MiniSeriesDTO{
						Wins:     0,
						Losses:   1,
						Target:   2,
						Progress: "LNN",
					},
					FreshBlood: false,
					Veteran:    false,
					Inactive:   false,
				},
			},
		}},
	}
	for _, test := range tests {
		jsonData, err := ioutil.ReadFile("test_data/league/leagues/" + test.leagueID + ".json")
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
		got, err := c.Leagues(ctx, test.leagueID)
		if err != nil {
			t.Fatalf("query returned with error: %v", err)
		}
		if LeagueListEqual(got, test.want) {
			t.Errorf("%v, want %v", got, test.want)
		}
	}
}

func TestEntries(t *testing.T) {
	limiter := rate.NewLimiter(20, 1)
	doer := http.DefaultClient
	c := client{
		HTTPClient: doer,
		Limiter:    limiter,
	}
	var tests = []struct {
		queue    QueueType
		tier     Tier
		division Division
		want     []LeagueEntryDTO
	}{
		{FlexSR, Gold, II, []LeagueEntryDTO{
			LeagueEntryDTO{
				SummonerName: "Bennouze",
				SummonerID:   "lvCSK82TNDHEiPzg_ib4sJkq2naWVw1FpggLlZQeFh9za-7g",
				LeagueID:     "",
				QueueType:    "RANKED_FLEX_SR",
				Tier:         "GOLD",
				Rank:         "II",
				LeaguePoints: 40,
				Wins:         105,
				Losses:       101,
				HotStreak:    false,
				MiniSeries: MiniSeriesDTO{
					Wins:     0,
					Losses:   0,
					Target:   0,
					Progress: "",
				},
				FreshBlood: false,
				Veteran:    true,
				Inactive:   false,
			},
			LeagueEntryDTO{
				SummonerName: "foop2legend",
				SummonerID:   "2D-hcgk5znPdkkjz1mKhkagH5mhDGsiKYn5oky-Tkhiruyul",
				LeagueID:     "",
				QueueType:    "RANKED_FLEX_SR",
				Tier:         "GOLD",
				Rank:         "II",
				LeaguePoints: 100,
				Wins:         112,
				Losses:       93,
				HotStreak:    false,
				MiniSeries: MiniSeriesDTO{
					Wins:     1,
					Losses:   1,
					Target:   2,
					Progress: "WLN",
				},
				FreshBlood: false,
				Veteran:    true,
				Inactive:   false,
			},
			LeagueEntryDTO{
				SummonerName: "M 4 R X M K X V",
				SummonerID:   "7bY6HkVIai_QCkNtQgbz4fl6FbEaAlDxSfZsGEjeXE1bHADc",
				LeagueID:     "",
				QueueType:    "RANKED_FLEX_SR",
				Tier:         "GOLD",
				Rank:         "II",
				LeaguePoints: 13,
				Wins:         10,
				Losses:       6,
				HotStreak:    false,
				MiniSeries: MiniSeriesDTO{
					Wins:     0,
					Losses:   0,
					Target:   0,
					Progress: "",
				},
				FreshBlood: false,
				Veteran:    false,
				Inactive:   false,
			},
		}},
		{FlexTT, Diamond, I, []LeagueEntryDTO{
			LeagueEntryDTO{
				SummonerName: "GOAT Jungler",
				SummonerID:   "fIN2VYHD6XgVsc4XdYh5kXt4tlqpC3hZtj7nW5t4c3Edtjcm",
				LeagueID:     "",
				QueueType:    "RANKED_FLEX_TT",
				Tier:         "DIAMOND",
				Rank:         "I",
				LeaguePoints: 75,
				Wins:         136,
				Losses:       122,
				HotStreak:    false,
				MiniSeries: MiniSeriesDTO{
					Wins:     0,
					Losses:   0,
					Target:   0,
					Progress: "",
				},
				FreshBlood: false,
				Veteran:    false,
				Inactive:   false,
			},
			LeagueEntryDTO{
				SummonerName: "26 11 2000",
				SummonerID:   "tz6T8-BlWlXR0p-vQw4idWTRwmXUpJt72q-SeqPJv8pYlDl4",
				LeagueID:     "",
				QueueType:    "RANKED_FLEX_TT",
				Tier:         "DIAMOND",
				Rank:         "I",
				LeaguePoints: 31,
				Wins:         27,
				Losses:       30,
				HotStreak:    false,
				MiniSeries: MiniSeriesDTO{
					Wins:     0,
					Losses:   0,
					Target:   0,
					Progress: "",
				},
				FreshBlood: false,
				Veteran:    false,
				Inactive:   false,
			},
			LeagueEntryDTO{
				SummonerName: "Brave Nugget",
				SummonerID:   "DhgQd99LwAiIGTi3lm1jRR-abDC3M7wkjzgaMPSoXspIXBis",
				LeagueID:     "",
				QueueType:    "RANKED_FLEX_TT",
				Tier:         "DIAMOND",
				Rank:         "I",
				LeaguePoints: 0,
				Wins:         12,
				Losses:       7,
				HotStreak:    false,
				MiniSeries: MiniSeriesDTO{
					Wins:     0,
					Losses:   0,
					Target:   0,
					Progress: "",
				},
				FreshBlood: false,
				Veteran:    false,
				Inactive:   false,
			},
		}},
		{Solo, Iron, IV, []LeagueEntryDTO{
			LeagueEntryDTO{
				SummonerName: "dav1d 370",
				SummonerID:   "AlTBwYLspM7n0dp5apF1bSGvWq9LePEu-2LxTvV6aNpmVf6q",
				LeagueID:     "",
				QueueType:    "RANKED_SOLO_5x5",
				Tier:         "IRON",
				Rank:         "IV",
				LeaguePoints: 75,
				Wins:         10,
				Losses:       16,
				HotStreak:    false,
				MiniSeries: MiniSeriesDTO{
					Wins:     0,
					Losses:   0,
					Target:   0,
					Progress: "",
				},
				FreshBlood: false,
				Veteran:    false,
				Inactive:   false,
			},
			LeagueEntryDTO{
				SummonerName: "PTRusherPT",
				SummonerID:   "ng2FNb7hqYqinqCB1HhJyP8CieUqxb7aC6QA1sw4gUO_5KWy",
				LeagueID:     "",
				QueueType:    "RANKED_SOLO_5x5",
				Tier:         "IRON",
				Rank:         "IV",
				LeaguePoints: 93,
				Wins:         4,
				Losses:       10,
				HotStreak:    false,
				MiniSeries: MiniSeriesDTO{
					Wins:     0,
					Losses:   0,
					Target:   0,
					Progress: "",
				},
				FreshBlood: false,
				Veteran:    false,
				Inactive:   false,
			},
			LeagueEntryDTO{
				SummonerName: "sami1999co",
				SummonerID:   "95uI4FWETsscOdNfdDTtl57KCNL530jV6fTIlUJ-8Ev6W3xx",
				LeagueID:     "",
				QueueType:    "RANKED_SOLO_5x5",
				Tier:         "IRON",
				Rank:         "IV",
				LeaguePoints: 95,
				Wins:         19,
				Losses:       21,
				HotStreak:    false,
				MiniSeries: MiniSeriesDTO{
					Wins:     0,
					Losses:   0,
					Target:   0,
					Progress: "",
				},
				FreshBlood: false,
				Veteran:    false,
				Inactive:   false,
			},
		}},
	}
	for _, test := range tests {
		jsonData, err := ioutil.ReadFile("test_data/league/entries/" + string(test.queue) + string(test.tier) + string(test.division) + ".json")
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
		got, err := c.Entries(ctx, test.queue, test.tier, test.division)
		if err != nil {
			t.Fatalf("query returned with error: %v", err)
		}
		if !LeagueEntrySliceEqual(got, test.want) {
			t.Errorf("%v, want %v", got, test.want)
		}
	}
}

func TestEntriesBySummoner(t *testing.T) {
	limiter := rate.NewLimiter(20, 1)
	doer := http.DefaultClient
	c := client{
		HTTPClient: doer,
		Limiter:    limiter,
	}
	var tests = []struct {
		summonerID string
		want       []LeagueEntryDTO
	}{
		{"_JEwwuOSAIZuxsZYp1V3mObKjIWjle8eAZgXWfoksGns74U", []LeagueEntryDTO{
			LeagueEntryDTO{
				SummonerName: "Mikrogeo",
				SummonerID:   "_JEwwuOSAIZuxsZYp1V3mObKjIWjle8eAZgXWfoksGns74U",
				LeagueID:     "7459b630-6492-11e9-896f-c81f66dd0e0d",
				QueueType:    "RANKED_FLEX_SR",
				Tier:         "GOLD",
				Rank:         "IV",
				LeaguePoints: 40,
				Wins:         7,
				Losses:       11,
				HotStreak:    false,
				MiniSeries: MiniSeriesDTO{
					Wins:     0,
					Losses:   0,
					Target:   0,
					Progress: "",
				},
				FreshBlood: false,
				Veteran:    false,
				Inactive:   false,
			},
		}},
		{"hv02SPkEwbGjIpwBhULUWj3ksoEkdTq3R7SSdmV-BKGrjDU", []LeagueEntryDTO{
			LeagueEntryDTO{
				SummonerName: "ZkCat",
				SummonerID:   "hv02SPkEwbGjIpwBhULUWj3ksoEkdTq3R7SSdmV-BKGrjDU",
				LeagueID:     "7459b630-6492-11e9-896f-c81f66dd0e0d",
				QueueType:    "RANKED_FLEX_SR",
				Tier:         "GOLD",
				Rank:         "III",
				LeaguePoints: 46,
				Wins:         8,
				Losses:       11,
				HotStreak:    false,
				MiniSeries: MiniSeriesDTO{
					Wins:     0,
					Losses:   0,
					Target:   0,
					Progress: "",
				},
				FreshBlood: false,
				Veteran:    false,
				Inactive:   false,
			},
			LeagueEntryDTO{
				SummonerName: "ZkCat",
				SummonerID:   "hv02SPkEwbGjIpwBhULUWj3ksoEkdTq3R7SSdmV-BKGrjDU",
				LeagueID:     "e4c72470-208b-11e9-9883-c81f66dd2a8f",
				QueueType:    "RANKED_SOLO_5x5",
				Tier:         "PLATINUM",
				Rank:         "II",
				LeaguePoints: 21,
				Wins:         18,
				Losses:       24,
				HotStreak:    false,
				MiniSeries: MiniSeriesDTO{
					Wins:     0,
					Losses:   0,
					Target:   0,
					Progress: "",
				},
				FreshBlood: false,
				Veteran:    false,
				Inactive:   false,
			},
		}},
	}
	for _, test := range tests {
		jsonData, err := ioutil.ReadFile("test_data/league/by_summoner/" + test.summonerID + ".json")
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
		got, err := c.EntriesBySummoner(ctx, test.summonerID)
		if err != nil {
			t.Fatalf("query returned with error: %v", err)
		}
		if !LeagueEntrySliceEqual(got, test.want) {
			t.Errorf("%v, want %v", got, test.want)
		}
	}
}

func LeagueListEqual(a, b LeagueListDTO) bool {
	if a.LeagueID != b.LeagueID && a.Name != b.Name && a.Queue != b.Queue && a.Tier != b.Tier {
		return false
	}
	for i, v := range a.Entries {
		if v != b.Entries[i] {
			return false
		}
	}
	return true
}

func LeagueEntrySliceEqual(a, b []LeagueEntryDTO) bool {
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
