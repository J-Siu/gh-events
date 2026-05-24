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
	"strconv"
	"strings"
	"time"

	"github.com/J-Siu/gh-events/global"
	"github.com/J-Siu/gh-events/schema"
	"github.com/juju/ansiterm"
)

// processed event info from GitHub event
type EventInfo struct {
	Skipped       bool   // whether event is fully handled by New()
	StrAction     string // custom action name, eg. created -> commented
	StrLogin      string // original event actor login
	StrRepo       string // original event repo name
	StrTime       string // original time string from GitHub in UTC
	StrTxt        string // usually title of issue and PR, depends on event type
	StrTxtPrefix  string // prefix for StrTxt, eg. "Issue#", "PR#", etc.
	StrType       string // original event type
	StrTypeAction string // original action name from GitHub, eg. created, labeled, etc.
	StrUrl        string // url for PR, issue, comment, publish, depends on event type, default to repo url
}

func (t *EventInfo) New(event *schema.Event) *EventInfo {
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
		if event.Payload.Description != nil {
			t.StrTxt = *event.Payload.Description
		}
		t.StrAction = "created"
		t.StrUrl += "/tree/" + *event.Payload.Ref
	case "ForkEvent": // nothing to do
		t.StrUrl = *event.Payload.Forkee.HtmlUrl
	case "IssueCommentEvent":
		switch *event.Payload.Action {
		case "created":
			t.StrAction = "commented"
		}
		if event.Payload.Issue.PullRequest == nil {
			t.StrTxtPrefix = "Issue#" // prefix for Issue
		} else {
			t.StrTxtPrefix = "PR#" // prefix for PR
		}
		t.StrTxtPrefix += strconv.FormatInt(*event.Payload.Issue.Number, 10)
		t.StrTxt = *event.Payload.Issue.Title
		t.StrUrl = *event.Payload.Comment.HtmlUrl
	case "IssuesEvent":
		t.StrTxtPrefix = "Issue#" + strconv.FormatInt(*event.Payload.Issue.Number, 10)
		switch *event.Payload.Action {
		case "labeled", "unlabeled":
			t.StrTxt = "label: " + event.Payload.Labels.Names()
		default:
			t.StrTxt = *event.Payload.Issue.Title
		}
		t.StrUrl = *event.Payload.Issue.HtmlUrl
	case "PullRequestEvent":
		t.StrTxtPrefix = "PR#" + strconv.FormatInt(*event.Payload.PR.Number, 10)
		switch *event.Payload.Action {
		case "labeled", "unlabeled":
			t.StrTxt = "label: " + event.Payload.Labels.Names()
		}
		t.StrUrl += "/pull/" + strconv.FormatInt(*event.Payload.PR.Number, 10)
	case "PullRequestReviewCommentEvent":
		switch *event.Payload.Action {
		case "created":
			t.StrAction = "commented"
		}
		t.StrTxtPrefix = "PR#" + strconv.FormatInt(*event.Payload.PR.Number, 10)
		t.StrUrl = *event.Payload.Comment.HtmlUrl
	case "PullRequestReviewEvent":
		switch *event.Payload.Action {
		case "created":
			t.StrAction = "reviewed"
		}
		t.StrTxtPrefix = "PR#" + strconv.FormatInt(*event.Payload.PR.Number, 10)
		t.StrUrl = *event.Payload.Review.HtmlUrl
	case "ReleaseEvent":
		t.StrUrl = *event.Payload.Release.HtmlUrl
	case "WatchEvent":
		t.StrAction = "starred"
	default:
		t.Skipped = true
	}

	return t
}

// list of EventInfo, with filter and string functions
type EventInfos struct {
	*OutputProperties
	List []*EventInfo
}

func (t *EventInfos) New(op *OutputProperties, events *[]schema.Event) *EventInfos {
	t.OutputProperties = op
	for _, event := range *events {
		var info EventInfo
		info.New(&event)
		if !t.Has(&info) {
			t.List = append(t.List, &info)
		}
	}
	return t
}

func (t *EventInfos) Filter() *EventInfos {
	n := new(EventInfos)
	n.OutputProperties = t.OutputProperties
	for _, e := range t.List {
		if len(t.Filters) > 0 && MatchFilter(t.Filters, e.StrAction, e.StrTypeAction, e.StrType) {
			continue
		}
		n.List = append(n.List, e)
	}
	return n
}

func (t *EventInfos) Has(info *EventInfo) bool {
	for _, i := range t.List {
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

func (t *EventInfos) String() string {
	var (
		strBuilder strings.Builder
		tabWriter  = ansiterm.NewTabWriter(&strBuilder, 1, 1, 1, ' ', 0)
	)

	for _, info := range t.List {
		if !t.All && info.Skipped {
			continue
		}
		strTxt := info.StrTxt
		if t.TimeLocal {
			if utc, err := time.Parse(global.STR_TIME_FORMAT_UTC, info.StrTime); err == nil {
				info.StrTime = utc.Local().Format(global.STR_TIME_FORMAT_LOCAL)
			}
		}
		if !t.TimeUTC && !t.TimeLocal {
			info.StrTime = ""
		}
		if info.Skipped {
			info.StrAction = global.STR_SKIPPED
		}
		if t.ShowType {
			if info.StrTypeAction == "" {
				info.StrAction += "\t(" + info.StrType + ")"
			} else {
				info.StrAction += "\t(" + info.StrType + ":" + info.StrTypeAction + ")"
			}
		}
		if t.ShowUrl {
			info.StrLogin = global.URL_GITHUB + "/" + info.StrLogin
			info.StrRepo = info.StrUrl
		} else {
			if len(info.StrTxtPrefix) > 0 {
				strTxt = info.StrTxtPrefix + " " + strTxt
			}
		}
		fmt.Fprintln(tabWriter, strings.TrimSpace(strings.Join([]string{info.StrTime, info.StrLogin, info.StrAction, info.StrRepo + " " + strTxt}, "\t")))
	}
	tabWriter.Flush()
	return strBuilder.String()
}
