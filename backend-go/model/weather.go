package model

type Weather struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Description string  `json:"weather"`
	Humidity    int     `json:"humidity,omitempty"`
	WindSpeed   float64 `json:"wind_speed,omitempty"`
	Icon        string  `json:"icon,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
