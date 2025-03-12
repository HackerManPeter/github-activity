package internal

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/HackerManPeter/github-activity/config"
)

var Client = &http.Client{
	Timeout: 30 * time.Second,
}

type GithubResponse struct {
	ID    string      `json:"id"`
	Type  GithubEvent `json:"type"`
	Actor struct {
		ID           int    `json:"id"`
		Login        string `json:"login"`
		DisplayLogin string `json:"display_login"`
		GravatarID   string `json:"gravatar_id"`
		URL          string `json:"url"`
		AvatarURL    string `json:"avatar_url"`
	} `json:"actor"`
	Repo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
	Payload   map[string]any `json:"payload"`
	Public    bool           `json:"public"`
	CreatedAt time.Time      `json:"created_at"`
	Org       struct {
		ID         int    `json:"id"`
		Login      string `json:"login"`
		GravatarID string `json:"gravatar_id"`
		URL        string `json:"url"`
		AvatarURL  string `json:"avatar_url"`
	} `json:"org"`
}

func BuildRequest(ctx context.Context, config *config.Config, username string) *http.Request {
	endpoint := fmt.Sprintf("https://api.github.com/users/%s/events?", username)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		panic("unable to make request")
	}

	authToken := fmt.Sprintf("Bearer %v", config.GithubToken)

	req.Header.Add("Authorization", authToken)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Add("Accept", "application/vnd.github+json")
	return req

}
