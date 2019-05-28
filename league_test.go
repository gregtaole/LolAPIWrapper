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
	t.Errorf("not implemented\n")
}

func TestMasterLeagueByQueue(t *testing.T) {
	t.Errorf("not implemented\n")
}

func TestLeagues(t *testing.T) {
	t.Errorf("not implemented\n")
}

func TestEntries(t *testing.T) {
	t.Errorf("not implemented\n")
}

func TestEntriesBySummoner(t *testing.T) {
	t.Errorf("not implemented\n")
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
