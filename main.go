package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/HackerManPeter/github-activity/config"
	"github.com/HackerManPeter/github-activity/internal"
)

func receiveArguments() string {
	if len(os.Args) != 2 {
		log.Fatal("Please pass in the username only")
		os.Exit(1)
	}
	return os.Args[1]
}

var conf = config.New()

func main() {
	ctx := context.Background()

	// get user input from cmd
	username := receiveArguments()

	// fetch data from github
	req := internal.BuildRequest(ctx, conf, username)
	resp, err := internal.Client.Do(req)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("Unable to reach github ~ Status Code: %v", resp.StatusCode)
	}

	// display data
	var githubResponse []internal.GithubResponse

	var response []string = []string{}

	err = json.NewDecoder(resp.Body).Decode(&githubResponse)
	if err != nil {
		fmt.Printf("%v", err)
	}

	for _, data := range githubResponse {

		switch data.Type {
		case internal.CreateEvent:
			var p internal.CreateEventStruct
			if err := internal.MapToStruct(data.Payload, &p); err != nil {
				log.Printf("Error parsing create event: %v", err)
				continue
			}

			response = append(response, data.FormatCreateEvent(p))

		case internal.DeleteEvent:
			var p internal.DeleteEventStruct
			if err := internal.MapToStruct(data.Payload, &p); err != nil {
				log.Printf("Error parsing delete event: %v", err)
				continue
			}

			response = append(response, data.FormatDeleteEvent(p))

		case internal.PushEvent:
			var p internal.PushEventStruct
			if err := internal.MapToStruct(data.Payload, &p); err != nil {
				log.Printf("Error parsing push event: %v", err)
				continue
			}

			response = append(response, data.FormatPushEvent(p))

		case internal.ForkEvent:
			var p internal.ForkEventStruct
			if err := internal.MapToStruct(data.Payload, &p); err != nil {
				log.Printf("Error parsing fork event: %v", err)
				continue
			}

			response = append(response, data.FormatForkEvent(p))

		case internal.GollumEvent:

		case internal.IssueCommentEvent:
			var p internal.IssueCommentEventStruct
			if err := internal.MapToStruct(data.Payload, &p); err != nil {
				log.Printf("Error parsing issue comment event: %v", err)
				continue
			}

			response = append(response, data.FormatIssueCommentEvent(p))

		case internal.IssuesEvent:
			var p internal.IssuesEventStruct
			if err := internal.MapToStruct(data.Payload, &p); err != nil {
				log.Printf("Error parsing issue event: %v", err)
				continue
			}

			response = append(response, data.FormatIssuesEvent(p))

		case internal.MemberEvent:
			var p internal.MemberEventStruct
			if err := internal.MapToStruct(data.Payload, &p); err != nil {
				log.Printf("Error parsing member event: %v", err)
				continue
			}

			response = append(response, data.FormatMemberEvent(p))

		case internal.PublicEvent:
			response = append(response, data.FormatPublicEvent())

		case internal.PullRequestEvent:
			var p internal.PullRequestEventStruct
			if err := internal.MapToStruct(data.Payload, &p); err != nil {
				log.Printf("Error parsing pull request event: %v", err)
				continue
			}

			response = append(response, data.FormatPullRequestEvent(p))

		}

	}

	for _, event := range response {
		fmt.Println(event)
	}

}
