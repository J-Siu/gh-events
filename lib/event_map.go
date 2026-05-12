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
	"github.com/J-Siu/go-helper/v2/strany"
)

// Array of maps return by client.Get()
type EventMaps struct {
	*OutputProperties
	List []any
}

func (t *EventMaps) New(op *OutputProperties, maps *[]any) *EventMaps {
	t.OutputProperties = op
	t.List = *maps
	return t
}

func (t *EventMaps) Filter() *EventMaps {
	n := new(EventMaps)
	n.OutputProperties = t.OutputProperties
	n.List = []any{}
	for _, e := range t.List {
		var (
			strAction string
			strType   string
		)
		strType, _ = e.(map[string]any)["type"].(string)
		strAction, _ = e.(map[string]any)["payload"].(map[string]any)["action"].(string)
		if len(t.Filters) > 0 && MatchFilter(t.Filters, strAction, "", strType) {
			continue
		}
		n.List = append(n.List, e)
	}
	return n
}

func (t *EventMaps) String() string { return *strany.String(t) }
