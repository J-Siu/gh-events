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

type Payload struct {
	*PayloadProperties
	PayloadPropertiesShared
}

type PayloadProperties struct {
	*CreateEvent
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
	Action     *string             `json:"action,omitempty"`       // ALL Events
	Assignee   any                 `json:"assignee,omitempty"`     // PullRequestEvent
	Assignees  any                 `json:"assignees,omitempty"`    // PullRequestEvent
	Comment    *Comment            `json:"comment,omitempty"`      // IssueCommentEvent
	Issue      *Issue              `json:"issue,omitempty"`        // IssueCommentEvent, IssueEvent
	Label      *Label              `json:"label,omitempty"`        // PullRequestEvent
	Labels     *Labels             `json:"labels,omitempty"`       // PullRequestEvent
	PR         *PullRequestMinimal `json:"pull_request,omitempty"` // PullRequestEvent, PullRequestReviewCommentEvent, PullRequestReviewEvent
	Ref        *string             `json:"ref,omitempty"`          // CreateEvent
	RefType    *string             `json:"ref_type,omitempty"`     // CreateEvent
	PusherType *string             `json:"pusher_type,omitempty"`  // CreateEvent
	FullRef    *string             `json:"full_ref,omitempty"`     // CreateEvent
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
	// PR_M      *PullRequestMinimal `json:"pull_request,omitempty"`
	Number *int64 `json:"number,omitempty"`
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

type CreateEvent struct {
	// Ref         *string `json:"ref,omitempty"`
	// RefType     *string `json:"ref_type,omitempty"`
	// PusherType  *string `json:"pusher_type,omitempty"`
	// FullRef     *string `json:"full_ref,omitempty"`
	Description *string `json:"description,omitempty"`
}

type DeleteEvent struct {
	// Ref        *string `json:"ref,omitempty"`
	// RefType    *string `json:"ref_type,omitempty"`
	// PusherType *string `json:"pusher_type,omitempty"`
	// FullRef    *string `json:"full_ref,omitempty"`
}

type DiscussionEvent struct{}
type GollumEvent struct{}
type MemberEvent struct{}
type PublicEvent struct{}
