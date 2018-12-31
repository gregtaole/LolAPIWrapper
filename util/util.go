package util

import (
	"fmt"
	"net/http"
)

func GetResponse(APIKey string, region string, endpoint string, param string) (*http.Response, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://%v.api.riotgames.com/lol/%v/%v", region, endpoint, param)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request for %v: %v", url, err)
	}
	fmt.Println(req.URL)
	req.Header.Set("X-Riot-Token", APIKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not get %v: %v", req.URL, err)
	}
	return resp, nil
}
