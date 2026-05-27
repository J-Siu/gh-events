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

package cmd

import (
	"fmt"
	"os"

	"github.com/J-Siu/gh-events/global"
	"github.com/J-Siu/gh-events/lib"
	"github.com/J-Siu/gh-events/schema"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/spf13/cobra"
)

var op = lib.EventsProperties{}

var rootCmd = &cobra.Command{
	Use:     "gh-events",
	Short:   "List Github api 'users/<USER>/events' output.",
	Version: global.Version,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			client, err = api.DefaultRESTClient()
			endpoint    string
			events      lib.IEvents
			response    any
		)

		if err == nil {
			var actor schema.Actor
			if err = client.Get("user", &actor); err == nil {
				endpoint = "users/" + *actor.Login
				if global.Flag.Public {
					endpoint += "/received_events"
				} else {
					endpoint += "/events"
				}
				if global.Flag.Json {
					events = new(lib.EventMaps)
					op.Maps = new([]lib.EventMap)
					response = op.Maps
				} else {
					events = new(lib.EventInfos)
					op.Events = new([]schema.Event)
					response = op.Events
				}
				if err = client.Get(endpoint, response); err == nil {
					fmt.Print(events.New(&op).Filter())
				}
			}
		}

		if err != nil {
			fmt.Println(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&global.Flag.Json, "json", "j", false, "show json (all flags ignored except -f, -p)")
	rootCmd.Flags().BoolVarP(&global.Flag.Public, "public", "p", false, "show /received_events")
	rootCmd.Flags().BoolVarP(&op.All, "all", "a", false, "show skipped event")
	rootCmd.Flags().BoolVarP(&op.ShowType, "type", "t", false, "show event type")
	rootCmd.Flags().BoolVarP(&op.ShowUrl, "url", "u", false, "show full url")
	rootCmd.Flags().BoolVarP(&op.TimeLocal, "local-time", "l", false, "show create time(local)")
	rootCmd.Flags().BoolVarP(&op.TimeUTC, "create-time", "c", false, "show create time(UTC)")
	rootCmd.Flags().StringArrayVarP(&op.Filters, "filter", "f", []string{}, "show events by action, type")

	rootCmd.MarkFlagsMutuallyExclusive("create-time", "local-time")
}
