package actions

import (
	"encoding/json"
	"io"
	"net/http"
)

type ActivitiesInfoJSON struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Commits []struct {
			Message string `json:"message"`
		} `json:"commits"`
		Ref_type string `json:"ref_type"`
		Action   string `json:"action"`
	} `json:"payload"`
}

func FetchGitHubActivities(gitHubName string) ([]ActivitiesInfoJSON, error) {

	URL := "https://api.github.com/users/" + gitHubName + "/events"
	req, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(req.Body)

	var activitiesList []ActivitiesInfoJSON
	if err := json.NewDecoder(req.Body).Decode(&activitiesList); err != nil {
		panic(err)
	}
	return activitiesList, nil
}
