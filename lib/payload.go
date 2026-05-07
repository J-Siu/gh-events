package lib

type Payload struct {
	Action    *string          `json:"action,omitempty"`        // ALL Events
	Assignee  any              `json:"assignee,omitempty"`      // PullRequestEvent
	Assignees any              `json:"assignees,omitempty"`     // PullRequestEvent
	Before    *string          `json:"before,omitempty"`        // PushEvent
	Comment   *Comment         `json:"comment,omitempty"`       // IssueCommentEvent, PullRequestReviewCommentEvent
	Forkee    *PL_Forkee       `json:"forkee,omitempty"`        // ForkEvent
	Head      *string          `json:"head,omitempty"`          // PushEvent
	Id        *int64           `json:"id,omitempty"`            // PullRequestEvent
	Issue     *PL_IssueEvent   `json:"issue,omitempty"`         // IssueEvent
	Label     any              `json:"label,omitempty"`         // PullRequestEvent
	Labels    any              `json:"labels,omitempty"`        // PullRequestEvent
	Number    *int64           `json:"number,omitempty"`        // PullRequestEvent
	PR        *PR_Minimal      `json:"pull_request,omitempty"`  // PullRequestEvent, PullRequestReviewCommentEvent, PullRequestReviewEvent
	PushId    *int64           `json:"push_id,omitempty"`       // PushEvent
	Ref       *string          `json:"ref,omitempty"`           // PushEvent
	Release   *PL_ReleaseEvent `json:"release,omitempty"`       // ReleaseEvent
	RepoId    *int64           `json:"repository_id,omitempty"` // PushEvent
	Review    *Review          `json:"review,omitempty"`        // PullRequestReviewEvent
}

type PayloadProperties struct {
}

// ForkEvent Payload
type PL_Forkee struct {
	Forks_url *string `json:"forks_url,omitempty"`
	FullName  *string `json:"full_name,omitempty"`
	Name      *string `json:"name,omitempty"`
	Url       *string `json:"url,omitempty"`
}

// IssueEvent Payload
type PL_IssueEvent struct {
	Comments_url  *string          `json:"comments_url"`
	Events_url    *string          `json:"events_url"`
	HtmlUrl       *string          `json:"html_url"`
	Number        *int64           `json:"number,omitempty"`
	PullRequest   *PL_IssueEventPR `json:"pull_request,omitempty"` // IssueCommentEvent
	RepositoryUrl *string          `json:"repository_url"`
	Title         *string          `json:"title,omitempty"`
	Url           *string          `json:"url"`
}

type PL_IssueEventPR struct {
	HtmlUrl *string `json:"html_url"`
	Url     *string `json:"url"`
}

type PL_ReleaseEvent struct {
	HtmlUrl *string `json:"html_url"`
	TagName *string `json:"tag_name"`
	Url     *string `json:"url"`
}
