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

// ForkEvent Payload
type Forkee struct {
	Forks_url *string `json:"forks_url,omitempty"`
	FullName  *string `json:"full_name,omitempty"`
	HtmlUrl   *string `json:"html_url"`
	Name      *string `json:"name,omitempty"`
	Url       *string `json:"url,omitempty"`
}

// IssueEvent Payload
type Issue struct {
	Comments_url  *string           `json:"comments_url"`
	Events_url    *string           `json:"events_url"`
	HtmlUrl       *string           `json:"html_url"`
	Number        *int64            `json:"number,omitempty"`
	PullRequest   *IssuePullRequest `json:"pull_request,omitempty"`
	RepositoryUrl *string           `json:"repository_url"`
	Title         *string           `json:"title,omitempty"`
	Url           *string           `json:"url"`
}

type IssuePullRequest struct {
	HtmlUrl *string `json:"html_url"`
	Url     *string `json:"url"`
}

// PullRequestEvent, PullRequestReviewCommentEvent, PullRequestReviewEvent
type PullRequestMinimal struct {
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

type Release struct {
	HtmlUrl *string `json:"html_url"`
	TagName *string `json:"tag_name"`
	Url     *string `json:"url"`
}

// PullRequestReviewEvent
type Review struct {
	HtmlUrl *string `json:"html_url"`
}
