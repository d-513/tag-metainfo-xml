//    Copyright (C) 2021 dada513
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.
//
//    You should have received a copy of the GNU General Public License
//    along with this program.  If not, see <https://www.gnu.org/licenses/>.
package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage:", os.Args[0], "<user> <repo> <filename>")
		os.Exit(1)
	}
	var client *github.Client
	pat, hasPat := os.LookupEnv("PAT")
	if hasPat {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: pat},
		)
		tc := oauth2.NewClient(ctx, ts)

		client = github.NewClient(tc)
	} else {
		fmt.Println("warning: not running with a personal access token. Expect rate limits")
		client = github.NewClient(nil)
	}
	opt := github.ListOptions{}
	tags, _, err := client.Repositories.ListTags(context.Background(), os.Args[1], os.Args[2], &opt)
	HandleError(err)

	filename := os.Args[3]
	f, erra := ioutil.ReadFile(filename)
	HandleError(erra)
	t := &component{}
	xml.Unmarshal(f, t)
	t.Releases.Release = []Release{}
	for _, tag := range tags {
		opt2 := github.ListOptions{}
		commit, _, erra := client.Repositories.GetCommit(context.Background(), os.Args[1], os.Args[2], tag.GetCommit().GetSHA(), &opt2)
		HandleError(erra)
		t.Releases.Release = append(t.Releases.Release, Release{Version: tag.GetName(), Date: commit.GetCommit().GetAuthor().GetDate().Format("2006-01-02")})
	}
	out, errb := xml.MarshalIndent(t, "", "  ")
	out = []byte(xml.Header + string(out))
	fmt.Println(string(out))

	HandleError(errb)
	ioutil.WriteFile(filename, out, 0644)
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
