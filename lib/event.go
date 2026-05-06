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

package lib

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/J-Siu/gh-events/global"
	"github.com/J-Siu/go-helper/v2/str"
	"github.com/juju/ansiterm"
)

type Event struct {
	Actor     *Actor   `json:"actor"`
	CreatedAt *string  `json:"created_at"`
	ID        *string  `json:"id"`
	Org       *Actor   `json:"org,omitempty"`
	Payload   *Payload `json:"payload"`
	Public    *bool    `json:"public"`
	Repo      *Repo    `json:"repo"`
	Type      *string  `json:"type"`
}

type Actor struct {
	AvatarURL    *string `json:"avatar_url"`
	DisplayLogin *string `json:"display_login"`
	GravatarId   *string `json:"gravatar_id"`
	HtmlUrl      *string `json:"html_url"` // User, Simple User
	Id           *int64  `json:"id"`
	Login        *string `json:"login"`
	Url          *string `json:"url"`
}

type Repo struct {
	Id   *int64  `json:"id"`
	Name *string `json:"name"`
	Url  *string `json:"url"`
}

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

// PullRequestEvent, PullRequestReviewCommentEvent, PullRequestReviewEvent
type PR_Minimal struct {
	Base   any     `json:"base"`
	Head   any     `json:"head"`
	Id     *int64  `json:"id,omitempty"`
	Number *int64  `json:"number,omitempty"`
	Url    *string `json:"url"`
}

// IssueCommentEvent, PullRequestReviewCommentEvent
type Comment struct {
	HtmlUrl  *string `json:"html_url"`
	IssueUrl *string `json:"issue_url"`
}

// PullRequestReviewEvent
type Review struct {
	HtmlUrl *string `json:"html_url"`
}

type EventInfo struct {
	StrAction     string // custom action name, eg. created -> commented
	StrLogin      string
	StrRepo       string
	StrTime       string
	StrTxt        string
	StrType       string
	StrTypeAction string
	StrUrl        string // url for PR, issue, comment, publish, depends on event type, default to repo url
}

func (e *Event) GetInfo() (info EventInfo) {
	var (
		skip         bool
		actionFilter []string
	)

	info.StrLogin = *e.Actor.Login
	info.StrRepo = *e.Repo.Name
	info.StrTime = *e.CreatedAt
	info.StrType = *e.Type
	info.StrUrl = global.URL_GITHUB + "/" + info.StrRepo
	if e.Payload.Action != nil {
		info.StrAction = *e.Payload.Action
		info.StrTypeAction = *e.Payload.Action
	}

	switch *e.Type {
	case "ForkEvent":
		// nothing to do
	case "IssueCommentEvent":
		actionFilter = []string{"labeled"}
		if str.ArrayContains(&actionFilter, e.Payload.Action, false) {
			skip = true
		}
		switch *e.Payload.Action {
		case "created":
			info.StrAction = "commented"
		}
		if e.Payload.Issue.PullRequest == nil {
			info.StrTxt = "Issue" // prefix for Issue
		} else {
			info.StrTxt = "PR#" // prefix for PR
		}
		info.StrTxt = info.StrTxt + strconv.FormatInt(*e.Payload.Issue.Number, 10) + " " + *e.Payload.Issue.Title
		info.StrUrl = *e.Payload.Comment.HtmlUrl
	case "IssuesEvent":
		actionFilter = []string{"labeled"}
		if str.ArrayContains(&actionFilter, e.Payload.Action, false) {
			skip = true
		}
		info.StrTxt = "Issue" + strconv.FormatInt(*e.Payload.Issue.Number, 10) + " " + *e.Payload.Issue.Title
		info.StrUrl = *e.Payload.Issue.HtmlUrl
	case "PullRequestEvent":
		actionFilter = []string{"labeled"}
		if str.ArrayContains(&actionFilter, e.Payload.Action, false) {
			skip = true
		}
		info.StrUrl = info.StrRepo + "/pull/" + strconv.FormatInt(*e.Payload.PR.Number, 10)
	case "PullRequestReviewEvent":
		switch *e.Payload.Action {
		case "created":
			info.StrAction = "reviewed"
		}
		info.StrTxt = "PR#" + strconv.FormatInt(*e.Payload.PR.Number, 10)
		info.StrUrl = *e.Payload.Review.HtmlUrl
	case "PullRequestReviewCommentEvent":
		switch *e.Payload.Action {
		case "created":
			info.StrAction = "commented"
		}
		info.StrTxt = " PR#" + strconv.FormatInt(*e.Payload.PR.Number, 10)
		info.StrUrl = *e.Payload.Comment.HtmlUrl
	case "WatchEvent":
		info.StrAction = "starred"
	case "ReleaseEvent":
		info.StrUrl = *e.Payload.Release.HtmlUrl
	default:
		skip = true
	}

	if skip {
		info.StrAction = global.STR_NOT_HANDLED
	}

	return info
}

type Events []Event

func (t *Events) Print(all, showTime, showType, showUrl bool) {
	var (
		tabWriter = ansiterm.NewTabWriter(os.Stdout, 1, 1, 1, ' ', 0)
	)

	for _, e := range *t {
		info := e.GetInfo()
		if !all && strings.EqualFold(info.StrAction, global.STR_NOT_HANDLED) {
			continue
		}
		if !showTime {
			info.StrTime = ""
		}
		if showType {
			if info.StrTypeAction == "" {
				info.StrAction += "\t(" + info.StrType + ")"
			} else {
				info.StrAction += "\t(" + info.StrType + ":" + info.StrTypeAction + ")"
			}
		}
		if showUrl {
			info.StrLogin = global.URL_GITHUB + "/" + info.StrLogin
			info.StrRepo = info.StrUrl
		}
		fmt.Fprintln(tabWriter, strings.Join([]string{info.StrTime, info.StrLogin, info.StrAction, info.StrRepo + " " + info.StrTxt}, "\t"))
	}
	tabWriter.Flush()
}
