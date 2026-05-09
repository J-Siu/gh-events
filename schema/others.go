/*
Copyright © 2026 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package schema

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

type Label struct {
	Color       *string `json:"color,omitempty"`
	Default     *bool   `json:"default,omitempty"`
	Description *string `json:"description,omitempty"`
	Id          *int64  `json:"id,omitempty"`
	Name        *string `json:"name"`
}

type Labels []Label

// returns a comma-separated string of all labels' names
func (t *Labels) Names() (n string) {
	for i, label := range *t {
		if i > 0 {
			n += ", "
		}
		n += *label.Name
	}
	return n
}
