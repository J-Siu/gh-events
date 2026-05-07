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
	EventProperties
}

type EventProperties struct {
	Actor     *Actor   `json:"actor"`
	CreatedAt *string  `json:"created_at"`
	ID        *string  `json:"id"`
	Org       *Actor   `json:"org,omitempty"`
	Payload   *Payload `json:"payload"`
	Public    *bool    `json:"public"`
	Repo      *Repo    `json:"repo"`
	Type      *string  `json:"type"`
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
		info.StrTxt = "PR#" + strconv.FormatInt(*e.Payload.PR.Number, 10)
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
