package json

// API: https://oslobysykkel.no/apne-data/sanntid

type Headers struct {
	LastUpdated int `json:"last_updated"`
}

// station_information.json, SI

type StationInformation struct {
	Headers
	Data SIData `json:"data"`
}

type SIData struct {
	Stations []SIDataStation `json:"stations"`
}

type SIDataStation struct {
	StationId string `json:"station_id"`
	Name      string `json:"name"`
}

// station_status.json, SS

type StationStatus struct {
	Headers
	Data SSData `json:"data"`
}

type SSData struct {
	Stations []SSDataStation `json:"stations"`
}

type SSDataStation struct {
	StationId     string `json:"station_id"`
	NumBikesAvail int    `json:"num_bikes_available"`
	NumDocksAvail int    `json:"num_docks_available"`
}

func (s StationStatus) Station(id string) (v SSDataStation, found bool) {
	for _, v = range s.Data.Stations {
		if v.StationId == id {
			return v, true
		}
	}
	return v, false
}
