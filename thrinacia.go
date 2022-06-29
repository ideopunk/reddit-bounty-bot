package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getCampaigns() ([]Post, error) {
	resp, err := http.Get("https://www.viaprize.org/api/service/restv1/campaign/")

	if err != nil {
		return nil, fmt.Errorf("could not request campaigns: %w", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	campaigns := []*Campaign{}

	if err := json.Unmarshal(body, &campaigns); err != nil {
		return nil, fmt.Errorf("could not unmarshall Thrinacia json: %w", err)
	}

	filteredCampaigns, err := filterCampaigns(campaigns)

	if err != nil {
		return nil, fmt.Errorf("could not filter campaigns: %w", err)
	}

	println(len(filteredCampaigns))
	posts := make([]Post, 0)
	for _, camp := range filteredCampaigns {
		posts = append(posts, mapCampaign(camp))
	}
	return posts, nil
}

func filterCampaigns(campaigns []*Campaign) ([]*Campaign, error) {
	helperArr := campaigns
	campaigns = campaigns[:0]
	for _, camp := range helperArr {

		date, err := time.Parse("2006-01-02 15:04:05-07", camp.Starts)
		if err != nil {
			return nil, fmt.Errorf("could not parse campaign date %v: %w", camp.Starts, err)
		}

		yesterday := time.Now().Add(-24 * time.Hour * 140)

		// if this campaign starts in the last 24 hours, include it.
		if date.After(yesterday) {
			campaigns = append(campaigns, camp)
		}
	}

	return campaigns, nil

}

func mapCampaign(campaign *Campaign) Post {
	return Post{Title: campaign.Name + ": " + campaign.Blurb, URL: "https://www.viaprize.org/" + campaign.URIPaths[0].Path}
}
