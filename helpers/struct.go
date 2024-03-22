package helpers

type Status struct {
	Status U `json:"users"`
}

type U struct {
	Wind        int    `json:"wind"`
	Water       int    `json:"water"`
	WindStatus  string `json:"wind_status"`
	WaterStatus string `json:"water_status"`
}