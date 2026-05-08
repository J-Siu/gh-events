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

func (t *EventInfo) New(event *Event) *EventInfo {
	var (
		actionFilter []string
		skip         bool
	)

	t.StrLogin = *event.Actor.Login
	t.StrRepo = *event.Repo.Name
	t.StrTime = *event.CreatedAt
	t.StrType = *event.Type
	t.StrUrl = global.URL_GITHUB + "/" + t.StrRepo
	if event.Payload.Action != nil {
		t.StrAction = *event.Payload.Action
		t.StrTypeAction = *event.Payload.Action
	}

	switch *event.Type {
	case "CreateEvent":
		t.StrTxt = *event.Payload.Description
	case "ForkEvent": // nothing to do
		t.StrUrl = *event.Payload.Forkee.HtmlUrl
	case "IssueCommentEvent":
		actionFilter = []string{"labeled"}
		if str.ArrayContains(&actionFilter, event.Payload.Action, false) {
			skip = true
		}
		switch *event.Payload.Action {
		case "created":
			t.StrAction = "commented"
		}
		if event.Payload.Issue.PullRequest == nil {
			t.StrTxt = "Issue#" // prefix for Issue
		} else {
			t.StrTxt = "PR#" // prefix for PR
		}
		t.StrTxt += strconv.FormatInt(*event.Payload.Issue.Number, 10) + " " + *event.Payload.Issue.Title
		t.StrUrl = *event.Payload.Comment.HtmlUrl
	case "IssuesEvent":
		t.StrTxt = "Issue#" + strconv.FormatInt(*event.Payload.Issue.Number, 10)
		switch *event.Payload.Action {
		case "labeled":
			t.StrTxt += " label: " + event.Payload.Labels.String()
		default:
			t.StrTxt += " " + *event.Payload.Issue.Title
		}
		t.StrUrl = *event.Payload.Issue.HtmlUrl
	case "PullRequestEvent":
		t.StrTxt = "PR#" + strconv.FormatInt(*event.Payload.PR.Number, 10)
		switch *event.Payload.Action {
		case "labeled":
			t.StrTxt += " label: " + event.Payload.Labels.String()
		}
		t.StrUrl += "/pull/" + strconv.FormatInt(*event.Payload.PR.Number, 10)
	case "PullRequestReviewCommentEvent":
		switch *event.Payload.Action {
		case "created":
			t.StrAction = "commented"
		}
		t.StrTxt = "PR#" + strconv.FormatInt(*event.Payload.PR.Number, 10)
		t.StrUrl = *event.Payload.Comment.HtmlUrl
	case "PullRequestReviewEvent":
		switch *event.Payload.Action {
		case "created":
			t.StrAction = "reviewed"
		}
		t.StrTxt = "PR#" + strconv.FormatInt(*event.Payload.PR.Number, 10)
		t.StrUrl = *event.Payload.Review.HtmlUrl
	case "ReleaseEvent":
		t.StrUrl = *event.Payload.Release.HtmlUrl
	case "WatchEvent":
		t.StrAction = "starred"
	default:
		skip = true
	}

	if skip {
		t.StrAction = global.STR_SKIPPED
	}

	return t
}

type EventInfoList []*EventInfo

func (t *EventInfoList) New(events *Events) *EventInfoList {
	for _, event := range *events {
		var info EventInfo
		info.New(&event)
		if !t.Exist(&info) {
			*t = append(*t, &info)
		}
	}
	return t
}

func (t *EventInfoList) Print(all, showTime, showType, showUrl bool, filter []string) {
	var (
		tabWriter = ansiterm.NewTabWriter(os.Stdout, 1, 1, 1, ' ', 0)
	)

	for _, info := range *t {
		if len(filter) > 0 && MatchFilter(filter, info.StrAction, info.StrTypeAction, info.StrType) {
			continue
		}
		if !all && (info.StrAction == global.STR_SKIPPED) {
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

func (t *EventInfoList) Exist(info *EventInfo) bool {
	for _, i := range *t {
		if i.StrLogin == info.StrLogin &&
			i.StrTxt == info.StrTxt &&
			i.StrType == info.StrType &&
			i.StrTypeAction == info.StrTypeAction &&
			i.StrUrl == info.StrUrl {
			return true
		}
	}
	return false
}
