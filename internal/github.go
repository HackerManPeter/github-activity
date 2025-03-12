package internal

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type GithubEvent = string

const (
	CommitCommentEvent            GithubEvent = "CommitCommentEvent"
	CreateEvent                   GithubEvent = "CreateEvent"
	DeleteEvent                   GithubEvent = "DeleteEvent"
	ForkEvent                     GithubEvent = "ForkEvent"
	GollumEvent                   GithubEvent = "GollumEvent"
	IssueCommentEvent             GithubEvent = "IssueCommentEvent"
	IssuesEvent                   GithubEvent = "IssuesEvent"
	MemberEvent                   GithubEvent = "MemberEvent"
	PublicEvent                   GithubEvent = "PublicEvent"
	PullRequestEvent              GithubEvent = "PullRequestEvent"
	PullRequestReviewEvent        GithubEvent = "PullRequestReviewEvent"
	PullRequestReviewCommentEvent GithubEvent = "PullRequestReviewCommentEvent"
	PullRequestReviewThreadEvent  GithubEvent = "PullRequestReviewThreadEvent"
	PushEvent                     GithubEvent = "PushEvent"
	ReleaseEvent                  GithubEvent = "ReleaseEvent"
	SponsorshipEvent              GithubEvent = "SponsorshipEvent"
	WatchEvent                    GithubEvent = "WatchEvent"
)

type CommitCommentEventStruct struct {
	//
}

type CreateEventStruct struct {
	Ref          string `json:"ref"`
	RefType      string `json:"ref_type"`
	MasterBranch string `json:"master_branch"`
	Description  string `json:"description"`
	PusherType   string `json:"pusher_type"`
}

func (r *GithubResponse) FormatCreateEvent(p CreateEventStruct) string {
	return fmt.Sprintf("- Created a %v on %v called %v", p.RefType, r.Repo.Name, p.Ref)
}

type DeleteEventStruct struct {
	Ref     string `json:"ref"`
	RefType string `json:"ref_type"`
}

func (r *GithubResponse) FormatDeleteEvent(p DeleteEventStruct) string {
	return fmt.Sprintf("- Deleted a %v on %v called %v", p.RefType, r.Repo.Name, p.Ref)
}

type ForkEventStruct struct {
	//
}

func (r *GithubResponse) FormatForkEvent(p ForkEventStruct) string {
	return fmt.Sprintf("- You forked %v", r.Repo.Name)
}

type Issue struct {
	Number int `json:"number"`
}

type IssueCommentEventStruct struct {
	Action string `json:"action"`
	Issue  Issue  `json:"issue"`
}

func (r *GithubResponse) FormatIssueCommentEvent(p IssueCommentEventStruct) string {
	c := cases.Title(language.English)
	return fmt.Sprintf("- %v comment on issue number %v on %v", c.String(p.Action), p.Issue.Number, r.Repo.Name)
}

type IssuesEventStruct struct {
	Action string `json:"action"`
	Issue  Issue  `json:"issue"`
}

func (r *GithubResponse) FormatIssuesEvent(p IssuesEventStruct) string {
	c := cases.Title(language.English)
	return fmt.Sprintf("- %v issue number %v on %v", c.String(p.Action), p.Issue.Number, r.Repo.Name)

}

type MemberEventStruct struct {
	Action string `json:"action"`
}

func (r *GithubResponse) FormatMemberEvent(p MemberEventStruct) string {
	return fmt.Sprintf("- Member %v to %v", p.Action, r.Repo.Name)
}

func (r *GithubResponse) FormatPublicEvent() string {
	return fmt.Sprintf("- %v was made public", r.Repo.Name)
}

type PullRequest struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
}

type PullRequestEventStruct struct {
	Action      string         `json:"action"`
	Name        string         `json:"name"`
	Changes     map[string]any `json:"changes"`
	PullRequest PullRequest    `json:"pull_request"`
	Reason      string         `json:"reason"`
}

func (r *GithubResponse) FormatPullRequestEvent(p PullRequestEventStruct) string {
	c := cases.Title(language.English)
	return fmt.Sprintf("- %v pull request \"%v (#%v)\" on %v", c.String(p.Action), p.PullRequest.Title, p.PullRequest.Number, r.Repo.Name)
}

type PullRequestReviewEventStruct struct {
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
}

func (r *GithubResponse) FormatPullRequestReviewEvent(p PullRequestEventStruct) string {
	c := cases.Title(language.English)
	return fmt.Sprintf("- %v a review on pull request \"%v (#%v)\"", c.String(p.Action), p.PullRequest.Title, p.PullRequest.Number)
}

type PushEventStruct struct {
	Commits []any `json:"commits"`
}

func (r *GithubResponse) FormatPushEvent(p PushEventStruct) string {
	return fmt.Sprintf("- Pushed %v commits to %v", len(p.Commits), r.Repo.Name)
}
