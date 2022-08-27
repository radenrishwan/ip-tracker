package web

import "ip-tracker/model/entity"

type LocationResponse struct {
	IpAddress string `json:"ip_address"`
	Region    string `json:"region"`
	City      string `json:"city"`
	Timezone  string `json:"timezone"`
}

func NewLocationResponse(location entity.Location) LocationResponse {
	return LocationResponse{
		IpAddress: location.IpAddress,
		Region:    location.Region,
		City:      location.City,
		Timezone:  location.Timezone,
	}
}
