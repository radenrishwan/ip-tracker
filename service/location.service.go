package service

import (
	"github.com/ip2location/ip2location-go/v9"
	"ip-tracker/model/entity"
	"ip-tracker/model/web"
)

type LocationService interface {
	Find(ip string) web.CommonResponse[web.LocationResponse]
}

type locationService struct {
	*ip2location.DB
}

func NewLocationService(db *ip2location.DB) LocationService {
	return &locationService{DB: db}
}

func (service *locationService) Find(ip string) web.CommonResponse[web.LocationResponse] {
	result, err := service.DB.Get_all(ip)
	if err != nil {
		panic(err)
	}

	location := entity.Location{
		City:      result.City,
		Region:    result.Region,
		Timezone:  result.Timezone,
		IpAddress: ip,
	}

	return web.CommonResponse[web.LocationResponse]{
		Status: "Success",
		Data:   web.NewLocationResponse(location),
	}
}
