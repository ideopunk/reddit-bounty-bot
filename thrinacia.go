package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getCampaigns() ([]*Campaign, error) {
	resp, err := http.Get("https://community.thrinacia.com/api/service/restv1/campaign/")

	if err != nil {
		return nil, fmt.Errorf("could not request campaigns: %w", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	fmt.Printf("%v", string(body))
	campaigns := []*Campaign{}

	if err := json.Unmarshal(body, &campaigns); err != nil {
		return nil, fmt.Errorf("could not unmarshall Thrinacia json: %w", err)
	}

	if err := filterCampaigns(campaigns); err != nil {
		return nil, fmt.Errorf("could not filter campaigns: %w", err)
	}

	return campaigns, nil
}

func filterCampaigns(campaigns []*Campaign) error {
	helperArr := campaigns
	campaigns = campaigns[:0]

	for _, camp := range helperArr {

		date, err := time.Parse("2006-01-02", camp.Starts)

		if err != nil {
			return fmt.Errorf("could not parse campaign date %v: %w", camp.Starts, err)
		}

		yesterday := time.Now().Add(-24 * time.Hour)

		// if this campaign starts in the last 24 hours, include it.
		if date.After(yesterday) {
			campaigns = append(campaigns, camp)
		}
	}

	return nil

}
