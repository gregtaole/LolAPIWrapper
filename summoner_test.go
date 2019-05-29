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

func TestSummonerByAccount(t *testing.T) {
	limiter := rate.NewLimiter(20, 1)
	doer := http.DefaultClient
	c := client{
		HTTPClient: doer,
		Limiter:    limiter,
	}
	var tests = []struct {
		accountID string
		want      SummonerDTO
	}{
		{"2DZY2QJXEvm1zUVmC-dyE87EKcfT2n56thH4DGHnVKCpUTw", SummonerDTO{
			Name:          "Dinervoid",
			ID:            "eolRXteXeqUQpy6lVgGSGYMGs7ht9cWFkvAlNUtOQ_01UzQ",
			AccountID:     "2DZY2QJXEvm1zUVmC-dyE87EKcfT2n56thH4DGHnVKCpUTw",
			Puuid:         "XHmNrfsNaJ0rRpCU6zVQakbpunyetW-zCGwEomsm1_0-XN4vtj4agwEiZbLeAQBuKURspO7Z2ThVrw",
			SummonerLevel: 57,
			ProfileIconID: 3015,
			RevisionDate:  1550101518000,
		}},
		{"4C35G_gK2wDN6hKxwImQimh16bwVRxpDfpyBQHpELpsMYVRZQZctnAGC", SummonerDTO{
			Name:          "DinerNoob",
			ID:            "6RyT_m6Bn-iUd1n3xkOthI1BVpuLjWj8aUMD8nrJogfMpUrZ",
			AccountID:     "4C35G_gK2wDN6hKxwImQimh16bwVRxpDfpyBQHpELpsMYVRZQZctnAGC",
			Puuid:         "E8gDZC1-CAsCqZ1TahgDXLH8TzcjSVncpxRmANKyV8tHxXBbC0eQCLA6-_ANs3QhsV3jrU4bwmaHjQ",
			SummonerLevel: 46,
			ProfileIconID: 750,
			RevisionDate:  1558814852000,
		}},
		{"t6xTXi-H5y2kkt4rsUr4sLBK4c207TfEVOG77ONn5ZLcRQ", SummonerDTO{
			Name:          "ZkCat",
			ID:            "hv02SPkEwbGjIpwBhULUWj3ksoEkdTq3R7SSdmV-BKGrjDU",
			AccountID:     "t6xTXi-H5y2kkt4rsUr4sLBK4c207TfEVOG77ONn5ZLcRQ",
			Puuid:         "5ba_Qb_mKxxAg4t6drnR-TMl1iELiHu3ulzK4EmrzonujW8sQsm73Fa4p1f0KP7jqHAjrOCdHh2Kmw",
			SummonerLevel: 152,
			ProfileIconID: 3790,
			RevisionDate:  1558786340000,
		}},
		{"Yt_Pv9-mJlDI8j3j8vvDs0MOcKSkpFvhs3Gca8O9iOSWE9s", SummonerDTO{
			Name:          "Mikrogeo",
			ID:            "_JEwwuOSAIZuxsZYp1V3mObKjIWjle8eAZgXWfoksGns74U",
			AccountID:     "Yt_Pv9-mJlDI8j3j8vvDs0MOcKSkpFvhs3Gca8O9iOSWE9s",
			Puuid:         "KRfzFRAmuCjxd7GsylVVLZdWd6X15IREkxcDKR-nC5UtPWo6KnH5mcLFm5Ujk3CwYvIYtjiLnB9jLA",
			SummonerLevel: 115,
			ProfileIconID: 4022,
			RevisionDate:  1558301354000,
		}},
	}
	for _, test := range tests {
		jsonData, err := ioutil.ReadFile("test_data/summoner/by_account/" + test.accountID + ".json")
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
		got, err := c.SummonerByAccount(ctx, test.accountID)
		if err != nil {
			t.Fatalf("query returned with error: %v", err)
		}
		if got != test.want {
			t.Errorf("%v, want %v", got, test.want)
		}
	}
}

func TestSummonerByID(t *testing.T) {
	limiter := rate.NewLimiter(20, 1)
	doer := http.DefaultClient
	c := client{
		HTTPClient: doer,
		Limiter:    limiter,
	}
	var tests = []struct {
		ID   string
		want SummonerDTO
	}{
		{"eolRXteXeqUQpy6lVgGSGYMGs7ht9cWFkvAlNUtOQ_01UzQ", SummonerDTO{
			Name:          "Dinervoid",
			ID:            "eolRXteXeqUQpy6lVgGSGYMGs7ht9cWFkvAlNUtOQ_01UzQ",
			AccountID:     "2DZY2QJXEvm1zUVmC-dyE87EKcfT2n56thH4DGHnVKCpUTw",
			Puuid:         "XHmNrfsNaJ0rRpCU6zVQakbpunyetW-zCGwEomsm1_0-XN4vtj4agwEiZbLeAQBuKURspO7Z2ThVrw",
			SummonerLevel: 57,
			ProfileIconID: 3015,
			RevisionDate:  1550101518000,
		}},
		{"6RyT_m6Bn-iUd1n3xkOthI1BVpuLjWj8aUMD8nrJogfMpUrZ", SummonerDTO{
			Name:          "DinerNoob",
			ID:            "6RyT_m6Bn-iUd1n3xkOthI1BVpuLjWj8aUMD8nrJogfMpUrZ",
			AccountID:     "4C35G_gK2wDN6hKxwImQimh16bwVRxpDfpyBQHpELpsMYVRZQZctnAGC",
			Puuid:         "E8gDZC1-CAsCqZ1TahgDXLH8TzcjSVncpxRmANKyV8tHxXBbC0eQCLA6-_ANs3QhsV3jrU4bwmaHjQ",
			SummonerLevel: 46,
			ProfileIconID: 750,
			RevisionDate:  1558814852000,
		}},
		{"hv02SPkEwbGjIpwBhULUWj3ksoEkdTq3R7SSdmV-BKGrjDU", SummonerDTO{
			Name:          "ZkCat",
			ID:            "hv02SPkEwbGjIpwBhULUWj3ksoEkdTq3R7SSdmV-BKGrjDU",
			AccountID:     "t6xTXi-H5y2kkt4rsUr4sLBK4c207TfEVOG77ONn5ZLcRQ",
			Puuid:         "5ba_Qb_mKxxAg4t6drnR-TMl1iELiHu3ulzK4EmrzonujW8sQsm73Fa4p1f0KP7jqHAjrOCdHh2Kmw",
			SummonerLevel: 152,
			ProfileIconID: 3790,
			RevisionDate:  1558786340000,
		}},
		{"_JEwwuOSAIZuxsZYp1V3mObKjIWjle8eAZgXWfoksGns74U", SummonerDTO{
			Name:          "Mikrogeo",
			ID:            "_JEwwuOSAIZuxsZYp1V3mObKjIWjle8eAZgXWfoksGns74U",
			AccountID:     "Yt_Pv9-mJlDI8j3j8vvDs0MOcKSkpFvhs3Gca8O9iOSWE9s",
			Puuid:         "KRfzFRAmuCjxd7GsylVVLZdWd6X15IREkxcDKR-nC5UtPWo6KnH5mcLFm5Ujk3CwYvIYtjiLnB9jLA",
			SummonerLevel: 115,
			ProfileIconID: 4022,
			RevisionDate:  1558301354000,
		}},
	}
	for _, test := range tests {
		jsonData, err := ioutil.ReadFile("test_data/summoner/by_id/" + test.ID + ".json")
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
		got, err := c.SummonerByAccount(ctx, test.ID)
		if err != nil {
			t.Fatalf("query returned with error: %v", err)
		}
		if got != test.want {
			t.Errorf("%v, want %v", got, test.want)
		}
	}
}

func TestSummonerByName(t *testing.T) {
	limiter := rate.NewLimiter(20, 1)
	doer := http.DefaultClient
	c := client{
		HTTPClient: doer,
		Limiter:    limiter,
	}
	var tests = []struct {
		name string
		want SummonerDTO
	}{
		{"Dinervoid", SummonerDTO{
			Name:          "Dinervoid",
			ID:            "eolRXteXeqUQpy6lVgGSGYMGs7ht9cWFkvAlNUtOQ_01UzQ",
			AccountID:     "2DZY2QJXEvm1zUVmC-dyE87EKcfT2n56thH4DGHnVKCpUTw",
			Puuid:         "XHmNrfsNaJ0rRpCU6zVQakbpunyetW-zCGwEomsm1_0-XN4vtj4agwEiZbLeAQBuKURspO7Z2ThVrw",
			SummonerLevel: 57,
			ProfileIconID: 3015,
			RevisionDate:  1550101518000,
		}},
		{"DinerNoob", SummonerDTO{
			Name:          "DinerNoob",
			ID:            "6RyT_m6Bn-iUd1n3xkOthI1BVpuLjWj8aUMD8nrJogfMpUrZ",
			AccountID:     "4C35G_gK2wDN6hKxwImQimh16bwVRxpDfpyBQHpELpsMYVRZQZctnAGC",
			Puuid:         "E8gDZC1-CAsCqZ1TahgDXLH8TzcjSVncpxRmANKyV8tHxXBbC0eQCLA6-_ANs3QhsV3jrU4bwmaHjQ",
			SummonerLevel: 46,
			ProfileIconID: 750,
			RevisionDate:  1558814852000,
		}},
		{"ZkCat", SummonerDTO{
			Name:          "ZkCat",
			ID:            "hv02SPkEwbGjIpwBhULUWj3ksoEkdTq3R7SSdmV-BKGrjDU",
			AccountID:     "t6xTXi-H5y2kkt4rsUr4sLBK4c207TfEVOG77ONn5ZLcRQ",
			Puuid:         "5ba_Qb_mKxxAg4t6drnR-TMl1iELiHu3ulzK4EmrzonujW8sQsm73Fa4p1f0KP7jqHAjrOCdHh2Kmw",
			SummonerLevel: 152,
			ProfileIconID: 3790,
			RevisionDate:  1558786340000,
		}},
		{"Mikrogeo", SummonerDTO{
			Name:          "Mikrogeo",
			ID:            "_JEwwuOSAIZuxsZYp1V3mObKjIWjle8eAZgXWfoksGns74U",
			AccountID:     "Yt_Pv9-mJlDI8j3j8vvDs0MOcKSkpFvhs3Gca8O9iOSWE9s",
			Puuid:         "KRfzFRAmuCjxd7GsylVVLZdWd6X15IREkxcDKR-nC5UtPWo6KnH5mcLFm5Ujk3CwYvIYtjiLnB9jLA",
			SummonerLevel: 115,
			ProfileIconID: 4022,
			RevisionDate:  1558301354000,
		}},
	}
	for _, test := range tests {
		jsonData, err := ioutil.ReadFile("test_data/summoner/by_name/" + test.name + ".json")
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
		got, err := c.SummonerByAccount(ctx, test.name)
		if err != nil {
			t.Fatalf("query returned with error: %v", err)
		}
		if got != test.want {
			t.Errorf("%v, want %v", got, test.want)
		}
	}
}

func TestSummonerByPuuid(t *testing.T) {
	limiter := rate.NewLimiter(20, 1)
	doer := http.DefaultClient
	c := client{
		HTTPClient: doer,
		Limiter:    limiter,
	}
	var tests = []struct {
		puuid string
		want  SummonerDTO
	}{
		{"XHmNrfsNaJ0rRpCU6zVQakbpunyetW-zCGwEomsm1_0-XN4vtj4agwEiZbLeAQBuKURspO7Z2ThVrw", SummonerDTO{
			Name:          "Dinervoid",
			ID:            "eolRXteXeqUQpy6lVgGSGYMGs7ht9cWFkvAlNUtOQ_01UzQ",
			AccountID:     "2DZY2QJXEvm1zUVmC-dyE87EKcfT2n56thH4DGHnVKCpUTw",
			Puuid:         "XHmNrfsNaJ0rRpCU6zVQakbpunyetW-zCGwEomsm1_0-XN4vtj4agwEiZbLeAQBuKURspO7Z2ThVrw",
			SummonerLevel: 57,
			ProfileIconID: 3015,
			RevisionDate:  1550101518000,
		}},
		{"E8gDZC1-CAsCqZ1TahgDXLH8TzcjSVncpxRmANKyV8tHxXBbC0eQCLA6-_ANs3QhsV3jrU4bwmaHjQ", SummonerDTO{
			Name:          "DinerNoob",
			ID:            "6RyT_m6Bn-iUd1n3xkOthI1BVpuLjWj8aUMD8nrJogfMpUrZ",
			AccountID:     "4C35G_gK2wDN6hKxwImQimh16bwVRxpDfpyBQHpELpsMYVRZQZctnAGC",
			Puuid:         "E8gDZC1-CAsCqZ1TahgDXLH8TzcjSVncpxRmANKyV8tHxXBbC0eQCLA6-_ANs3QhsV3jrU4bwmaHjQ",
			SummonerLevel: 46,
			ProfileIconID: 750,
			RevisionDate:  1558814852000,
		}},
		{"5ba_Qb_mKxxAg4t6drnR-TMl1iELiHu3ulzK4EmrzonujW8sQsm73Fa4p1f0KP7jqHAjrOCdHh2Kmw", SummonerDTO{
			Name:          "ZkCat",
			ID:            "hv02SPkEwbGjIpwBhULUWj3ksoEkdTq3R7SSdmV-BKGrjDU",
			AccountID:     "t6xTXi-H5y2kkt4rsUr4sLBK4c207TfEVOG77ONn5ZLcRQ",
			Puuid:         "5ba_Qb_mKxxAg4t6drnR-TMl1iELiHu3ulzK4EmrzonujW8sQsm73Fa4p1f0KP7jqHAjrOCdHh2Kmw",
			SummonerLevel: 152,
			ProfileIconID: 3790,
			RevisionDate:  1558786340000,
		}},
		{"KRfzFRAmuCjxd7GsylVVLZdWd6X15IREkxcDKR-nC5UtPWo6KnH5mcLFm5Ujk3CwYvIYtjiLnB9jLA", SummonerDTO{
			Name:          "Mikrogeo",
			ID:            "_JEwwuOSAIZuxsZYp1V3mObKjIWjle8eAZgXWfoksGns74U",
			AccountID:     "Yt_Pv9-mJlDI8j3j8vvDs0MOcKSkpFvhs3Gca8O9iOSWE9s",
			Puuid:         "KRfzFRAmuCjxd7GsylVVLZdWd6X15IREkxcDKR-nC5UtPWo6KnH5mcLFm5Ujk3CwYvIYtjiLnB9jLA",
			SummonerLevel: 115,
			ProfileIconID: 4022,
			RevisionDate:  1558301354000,
		}},
	}
	for _, test := range tests {
		jsonData, err := ioutil.ReadFile("test_data/summoner/by_puuid/" + test.puuid + ".json")
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
		got, err := c.SummonerByAccount(ctx, test.puuid)
		if err != nil {
			t.Fatalf("query returned with error: %v", err)
		}
		if got != test.want {
			t.Errorf("%v, want %v", got, test.want)
		}
	}
}
