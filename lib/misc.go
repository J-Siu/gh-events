package lib

type Actor struct {
	AvatarURL    *string `json:"avatar_url"`
	DisplayLogin *string `json:"display_login"`
	GravatarId   *string `json:"gravatar_id"`
	HtmlUrl      *string `json:"html_url"` // User, Simple User
	Id           *int64  `json:"id"`
	Login        *string `json:"login"`
	Url          *string `json:"url"`
}

// IssueCommentEvent, PullRequestReviewCommentEvent
type Comment struct {
	HtmlUrl  *string `json:"html_url"`
	IssueUrl *string `json:"issue_url"`
}

// PullRequestEvent, PullRequestReviewCommentEvent, PullRequestReviewEvent
type PR_Minimal struct {
	Base   any     `json:"base"`
	Head   any     `json:"head"`
	Id     *int64  `json:"id,omitempty"`
	Number *int64  `json:"number,omitempty"`
	Url    *string `json:"url"`
}

type Repo struct {
	Id   *int64  `json:"id"`
	Name *string `json:"name"`
	Url  *string `json:"url"`
}

// PullRequestReviewEvent
type Review struct {
	HtmlUrl *string `json:"html_url"`
}
