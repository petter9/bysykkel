package main

import (
	"flag"
	"fmt"
	"os"
	"p/bysykkel/json"
)

var unicode = flag.Bool("u", false, "unicode")

func main() {
	flag.Parse()

	bike := "Bikes:"
	dock := "Docks:"
	if *unicode {
		bike = "🚲 "
		dock = "⭮ "
	}

	do := func(info json.SIDataStation, s json.SSDataStation, found bool) {
		if !found {
			fmt.Printf("%-25s <data not found>\n", info.Name)
			return
		}
		fmt.Printf("%-25s %s%-2d %s%-2d\n",
			info.Name, bike, s.NumBikesAvail, dock, s.NumDocksAvail)
	}

	err := json.Stations(do)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
