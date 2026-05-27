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

type EventMap any

// Array of maps return by client.Get()
type EventMaps struct {
	*EventsProperties
}

func (t *EventMaps) New(op *EventsProperties) IEvents {
	t.EventsProperties = op
	return t
}

func (t *EventMaps) Filter() IEvents {
	n := new([]EventMap)
	for _, eventMap := range *t.EventsProperties.Maps {
		var (
			strAction string
			strType   string
		)
		strType, _ = eventMap.(map[string]any)["type"].(string)
		strAction, _ = eventMap.(map[string]any)["payload"].(map[string]any)["action"].(string)
		if len(t.Filters) > 0 && MatchFilter(t.Filters, strAction, "", strType) {
			continue
		}
		*n = append(*n, eventMap)
	}
	t.Maps = n
	return t
}

func (t *EventMaps) String() string { return strany.Any(t.EventsProperties.Maps) + "\n" }
