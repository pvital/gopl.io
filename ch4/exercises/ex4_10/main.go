// Copyright © 2021 Paulo Vital
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.

// Exercise 4.10 from section 4.5 - JSON.
// Prints a table of GitHub issues matching the search terms, divided by last
// update age categories - less than a month ago, less than a year ago, more
// than year ago.

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	// Get Issues matching the search terms from cmd-line arguments.
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d issues in total.\n", result.TotalCount)

	// Map of the categories to print.
	ages :=  map[string][]*github.Issue {
		// issues created less than a month ago
		"month_ago":	make([]*github.Issue, 0, result.TotalCount),
		// issues created less than a year ago
		"year_ago":		make([]*github.Issue, 0, result.TotalCount),
		// issues created more than a year ago
		"more_year":	make([]*github.Issue, 0, result.TotalCount),
	}

	today := time.Now()

	// Divide the issues into the categories
	for _, item := range result.Items {
		diff := today.Sub(item.UpdatedAt)

		// A month has in average 730.5h
		hours_in_month := 730.5
		if diff.Hours() <= hours_in_month {
			ages["month_ago"] = append(ages["month_ago"], item)
		} else if diff.Hours() <= (hours_in_month * 12) {
			ages["year_ago"] = append(ages["year_ago"], item)
		} else {
			ages["more_year"] = append(ages["more_year"], item)
		}
	}

	// Print results
	fmt.Printf("\n")
	fmt.Printf("Issues last updated less than a month ago (%d):\n",
		len(ages["month_ago"]))
	for _, item := range ages["month_ago"] {
		fmt.Printf("\t#%-5d %.55s\n", item.Number, item.Title)
	}
	fmt.Printf("\n")
	fmt.Printf("Issues last updated less than a year ago: (%d):\n",
		len(ages["year_ago"]))
	for _, item := range ages["year_ago"] {
		fmt.Printf("\t#%-5d %.55s\n", item.Number, item.Title)
	}
	fmt.Printf("\n")
	fmt.Printf("Issues last updated more than year ago (%d):\n",
		len(ages["more_year"]))
	for _, item := range ages["more_year"] {
		fmt.Printf("\t#%-5d %.55s\n", item.Number, item.Title)
	}
}

/*
//!+textoutput
$ go build -o issues_by_age ch4/execises/ex4_10/main.go
$ ./issues_by_age repo:golang/go is:open json decoder
Found 62 issues in total.
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
