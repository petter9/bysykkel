// Package json implements functions that fetches and unmarshals json resources.
package json

import (
	"encoding/json"
	"p/bysykkel/http"
	"p/bysykkel/url"
	"sort"
)

// Stations gathers data about all the stations and lets the caller run a
// function on each entry in this set.
func Stations(do func(SIDataStation, SSDataStation, bool)) error {
	var info StationInformation

	data, err := http.Fetch(url.StationInformation)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &info)
	if err != nil {
		return err
	}

	var status StationStatus

	data, err = http.Fetch(url.StationStatus)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &status)
	if err != nil {
		return err
	}

	// Sort by name
	sort.Slice(info.Data.Stations, func(i, j int) bool {
		return info.Data.Stations[i].Name < info.Data.Stations[j].Name
	})

	for _, is := range info.Data.Stations {
		ss, found := status.Station(is.StationId)
		do(is, ss, found)
	}

	return nil
}
