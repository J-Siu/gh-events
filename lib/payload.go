package lib

type Payload struct {
	*PayloadProperties
	PayloadPropertiesShared
}

type PayloadProperties struct {
	*ForkEvent
	*IssueCommentEvent
	*IssuesEvent
	*PullRequestEvent
	*PullRequestReviewCommentEvent
	*PullRequestReviewEvent
	*ReleaseEvent
	*WatchEvent
}
type PayloadPropertiesShared struct {
	Action    *string             `json:"action,omitempty"`       // ALL Events
	Assignee  any                 `json:"assignee,omitempty"`     // PullRequestEvent
	Assignees any                 `json:"assignees,omitempty"`    // PullRequestEvent
	Comment   *Comment            `json:"comment,omitempty"`      // IssueCommentEvent
	Issue     *Issue              `json:"issue,omitempty"`        // IssueCommentEvent, IssueEvent
	Label     any                 `json:"label,omitempty"`        // PullRequestEvent
	Labels    any                 `json:"labels,omitempty"`       // PullRequestEvent
	PR        *PullRequestMinimal `json:"pull_request,omitempty"` // PullRequestEvent, PullRequestReviewCommentEvent, PullRequestReviewEvent
}

type ForkEvent struct {
	Forkee *Forkee `json:"forkee,omitempty"` // ForkEvent
}

type IssueCommentEvent struct {
	// Action  *string  `json:"action,omitempty"`
	// Comment *Comment `json:"comment,omitempty"`
	// Issue   *Issue   `json:"issue,omitempty"`
}

type IssuesEvent struct {
	// Action    *string `json:"action,omitempty"`
	// Assignee  any     `json:"assignee,omitempty"`
	// Assignees any     `json:"assignees,omitempty"`
	// Issue     *Issue  `json:"issue,omitempty"`
	// Label     any     `json:"label,omitempty"`
	// Labels    any     `json:"labels,omitempty"`
}

type PullRequestEvent struct {
	// Action    *string             `json:"action,omitempty"`
	// Assignee  any                 `json:"assignee,omitempty"`
	// Assignees any                 `json:"assignees,omitempty"`
	// Label     any                 `json:"label,omitempty"`
	// Labels    any                 `json:"labels,omitempty"`
	// Number    *int64              `json:"number,omitempty"`
	// PR_M      *PullRequestMinimal `json:"pull_request,omitempty"`
}

type PullRequestReviewCommentEvent struct {
	// Action  *string             `json:"action,omitempty"`
	// Comment *Comment            `json:"comment,omitempty"`
	// PR_M    *PullRequestMinimal `json:"pull_request,omitempty"`
}

type PullRequestReviewEvent struct {
	// Action *string             `json:"action,omitempty"`
	// PR_M   *PullRequestMinimal `json:"pull_request,omitempty"`
	Review *Review `json:"review,omitempty"`
}

type PushEvent struct {
	Before *string `json:"before,omitempty"`
	Head   *string `json:"head,omitempty"`
	PushId *int64  `json:"push_id,omitempty"`
	Ref    *string `json:"ref,omitempty"`
	RepoId *int64  `json:"repository_id,omitempty"`
}

type ReleaseEvent struct {
	// Action  *string  `json:"action,omitempty"`
	Release *Release `json:"release,omitempty"`
}

type WatchEvent struct {
	// Action *string `json:"action,omitempty"`
}

type CreateEvent struct{}
type DeleteEvent struct{}
type DiscussionEvent struct{}
type GollumEvent struct{}
type MemberEvent struct{}
type PublicEvent struct{}
