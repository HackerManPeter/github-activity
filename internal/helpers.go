package internal

import (
	"encoding/json"
	"fmt"
	"log"
)

func MapToStruct[T any](m map[string]any, s *T) error {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, s)
	if err != nil {
		return err
	}

	return nil
}

func FormatOutput(u string, r []string) string {
	output := fmt.Sprintf("Recent GitHub Activity for %v:\n", u)

	for _, event := range r {
		output += fmt.Sprintln(event)
	}

	return output
}

func ParseGithubResponse(githubResponse []GithubResponse) []string {
	var response = []string{}
	for _, data := range githubResponse {

		switch data.Type {
		case CreateEvent:
			var p CreateEventStruct
			if err := MapToStruct(data.Payload, &p); err != nil {
				log.Printf("Error parsing create event: %v", err)
				continue
			}

			response = append(response, data.FormatCreateEvent(p))

		case DeleteEvent:
			var p DeleteEventStruct
			if err := MapToStruct(data.Payload, &p); err != nil {
				log.Printf("Error parsing delete event: %v", err)
				continue
			}

			response = append(response, data.FormatDeleteEvent(p))

		case PushEvent:
			var p PushEventStruct
			if err := MapToStruct(data.Payload, &p); err != nil {
				log.Printf("Error parsing push event: %v", err)
				continue
			}

			response = append(response, data.FormatPushEvent(p))

		case ForkEvent:
			var p ForkEventStruct
			if err := MapToStruct(data.Payload, &p); err != nil {
				log.Printf("Error parsing fork event: %v", err)
				continue
			}

			response = append(response, data.FormatForkEvent(p))

		case GollumEvent:

		case IssueCommentEvent:
			var p IssueCommentEventStruct
			if err := MapToStruct(data.Payload, &p); err != nil {
				log.Printf("Error parsing issue comment event: %v", err)
				continue
			}

			response = append(response, data.FormatIssueCommentEvent(p))

		case IssuesEvent:
			var p IssuesEventStruct
			if err := MapToStruct(data.Payload, &p); err != nil {
				log.Printf("Error parsing issue event: %v", err)
				continue
			}

			response = append(response, data.FormatIssuesEvent(p))

		case MemberEvent:
			var p MemberEventStruct
			if err := MapToStruct(data.Payload, &p); err != nil {
				log.Printf("Error parsing member event: %v", err)
				continue
			}

			response = append(response, data.FormatMemberEvent(p))

		case PublicEvent:
			response = append(response, data.FormatPublicEvent())

		case PullRequestEvent:
			var p PullRequestEventStruct
			if err := MapToStruct(data.Payload, &p); err != nil {
				log.Printf("Error parsing pull request event: %v", err)
				continue
			}

			response = append(response, data.FormatPullRequestEvent(p))

		}

	}

	return response
}
