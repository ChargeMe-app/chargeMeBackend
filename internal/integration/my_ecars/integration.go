package my_ecars

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/poorfrombabylon/chargeMeBackend/internal/config"
)

type Integration interface {
	GetAllStations(context.Context) (MyECarsStationsResponse, error)
	GetStationByID(context.Context, string) (MyECarsStationsResponse, error)
}

type MyECarsIntegration struct {
	baseUrl string
	key     string
}

func NewMyECarsIntegration(conf config.MyECars) Integration {
	return MyECarsIntegration{
		baseUrl: conf.BaseUrl,
		key:     conf.Key,
	}
}

func (m MyECarsIntegration) GetAllStations(_ context.Context) (MyECarsStationsResponse, error) {
	u, err := url.Parse(m.baseUrl)
	if err != nil {
		return MyECarsStationsResponse{}, err
	}

	data := url.Values{
		"key": {m.key},
	}

	resp, err := http.PostForm(u.String(), data)
	if err != nil {
		return MyECarsStationsResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return MyECarsStationsResponse{}, err
	}

	var result MyECarsStationsResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return MyECarsStationsResponse{}, err
	}

	return result, nil
}

func (m MyECarsIntegration) GetStationByID(_ context.Context, stationID string) (MyECarsStationsResponse, error) {
	u, err := url.Parse(m.baseUrl)
	if err != nil {
		return MyECarsStationsResponse{}, err
	}

	data := url.Values{
		"key": {m.key},
		"id":  {stationID},
	}

	resp, err := http.PostForm(u.String(), data)
	if err != nil {
		return MyECarsStationsResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return MyECarsStationsResponse{}, err
	}

	var result MyECarsStationsResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return MyECarsStationsResponse{}, err
	}

	return result, nil
}
