package main

import "time"

type Response []struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
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
	Payload struct {
		Action string `json:"action"`
		Issue  struct {
			URL           string `json:"url"`
			RepositoryURL string `json:"repository_url"`
			LabelsURL     string `json:"labels_url"`
			CommentsURL   string `json:"comments_url"`
			EventsURL     string `json:"events_url"`
			HTMLURL       string `json:"html_url"`
			ID            int    `json:"id"`
			NodeID        string `json:"node_id"`
			Number        int    `json:"number"`
			Title         string `json:"title"`
			User          struct {
				Login             string `json:"login"`
				ID                int    `json:"id"`
				NodeID            string `json:"node_id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				OrganizationsURL  string `json:"organizations_url"`
				ReposURL          string `json:"repos_url"`
				EventsURL         string `json:"events_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"user"`
			Labels            []interface{} `json:"labels"`
			State             string        `json:"state"`
			Locked            bool          `json:"locked"`
			Assignee          interface{}   `json:"assignee"`
			Assignees         []interface{} `json:"assignees"`
			Milestone         interface{}   `json:"milestone"`
			Comments          int           `json:"comments"`
			CreatedAt         time.Time     `json:"created_at"`
			UpdatedAt         time.Time     `json:"updated_at"`
			ClosedAt          interface{}   `json:"closed_at"`
			AuthorAssociation string        `json:"author_association"`
			ActiveLockReason  interface{}   `json:"active_lock_reason"`
			Body              string        `json:"body"`
			Reactions         struct {
				URL        string `json:"url"`
				TotalCount int    `json:"total_count"`
				Num1       int    `json:"+1"`
				Num10      int    `json:"-1"`
				Laugh      int    `json:"laugh"`
				Hooray     int    `json:"hooray"`
				Confused   int    `json:"confused"`
				Heart      int    `json:"heart"`
				Rocket     int    `json:"rocket"`
				Eyes       int    `json:"eyes"`
			} `json:"reactions"`
			TimelineURL           string      `json:"timeline_url"`
			PerformedViaGithubApp interface{} `json:"performed_via_github_app"`
			StateReason           interface{} `json:"state_reason"`
		} `json:"issue"`
		Comment struct {
			URL      string `json:"url"`
			HTMLURL  string `json:"html_url"`
			IssueURL string `json:"issue_url"`
			ID       int    `json:"id"`
			NodeID   string `json:"node_id"`
			User     struct {
				Login             string `json:"login"`
				ID                int    `json:"id"`
				NodeID            string `json:"node_id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				OrganizationsURL  string `json:"organizations_url"`
				ReposURL          string `json:"repos_url"`
				EventsURL         string `json:"events_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"user"`
			CreatedAt         time.Time `json:"created_at"`
			UpdatedAt         time.Time `json:"updated_at"`
			AuthorAssociation string    `json:"author_association"`
			Body              string    `json:"body"`
			Reactions         struct {
				URL        string `json:"url"`
				TotalCount int    `json:"total_count"`
				Num1       int    `json:"+1"`
				Num10      int    `json:"-1"`
				Laugh      int    `json:"laugh"`
				Hooray     int    `json:"hooray"`
				Confused   int    `json:"confused"`
				Heart      int    `json:"heart"`
				Rocket     int    `json:"rocket"`
				Eyes       int    `json:"eyes"`
			} `json:"reactions"`
			PerformedViaGithubApp interface{} `json:"performed_via_github_app"`
		} `json:"comment"`
	} `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
	Org       struct {
		ID         int    `json:"id"`
		Login      string `json:"login"`
		GravatarID string `json:"gravatar_id"`
		URL        string `json:"url"`
		AvatarURL  string `json:"avatar_url"`
	} `json:"org"`
}