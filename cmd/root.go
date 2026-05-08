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
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gh-events",
	Short:   "List Github api 'users/<USER>/received_events' output.",
	Version: global.Version,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			actor    lib.Actor
			endpoint string
		)

		// login user
		client, err := api.DefaultRESTClient()
		if err == nil {
			err = client.Get("user", &actor)
		}

		if err == nil {
			endpoint = "users/" + *actor.Login + "/received_events"
			if global.Flag.Public {
				endpoint += "/public"
			}

			if global.Flag.Json {
				var raw lib.EventsRaw
				if err = client.Get(endpoint, &raw); err == nil {
					raw.Print(global.Flag.Filter)
				}
			} else {
				var (
					events   lib.Events
					infoList lib.EventInfoList
				)
				if err = client.Get(endpoint, &events); err == nil {
					infoList.New(&events).Print(global.Flag.All, global.Flag.Time, global.Flag.Type, global.Flag.Url, global.Flag.Filter)
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
	rootCmd.Flags().BoolVarP(&global.Flag.All, "all", "a", false, "show skipped event")
	rootCmd.Flags().BoolVarP(&global.Flag.Json, "json", "j", false, "show raw json")
	rootCmd.Flags().BoolVarP(&global.Flag.Public, "public", "p", false, "show public events")
	rootCmd.Flags().BoolVarP(&global.Flag.Time, "create-time", "c", false, "show create time")
	rootCmd.Flags().BoolVarP(&global.Flag.Type, "type", "t", false, "show event type")
	rootCmd.Flags().BoolVarP(&global.Flag.Url, "url", "u", false, "show full url")
	rootCmd.Flags().StringArrayVarP(&global.Flag.Filter, "filter", "f", []string{}, "show events by action, type")
}
