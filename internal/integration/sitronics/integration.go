package sitronics

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/poorfrombabylon/chargeMeBackend/internal/config"
)

type Integration interface {
	GetAllStations(context.Context) (SitronicsMapInfo, error)
	GetStationByName(context.Context, string) (SitronicsStationInfo, error)
}

type sitronicsIntegration struct {
	baseUrl string
	client  *http.Client
}

func NewSitronicsIntegration(conf config.Sitronics) Integration {
	return sitronicsIntegration{
		baseUrl: conf.BaseUrl,
		client:  &http.Client{},
	}
}

func (s sitronicsIntegration) GetAllStations(_ context.Context) (SitronicsMapInfo, error) {
	u, err := url.Parse(s.baseUrl)
	if err != nil {
		return SitronicsMapInfo{}, err
	}

	u.Path = path.Join(u.Path, "CpApi", "GetMapDataSite")

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return SitronicsMapInfo{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return SitronicsMapInfo{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SitronicsMapInfo{}, err
	}

	var result SitronicsMapInfo
	err = json.Unmarshal(body, &result)
	if err != nil {
		return SitronicsMapInfo{}, err
	}

	return result, nil
}

func (s sitronicsIntegration) GetStationByName(_ context.Context, stationName string) (SitronicsStationInfo, error) {
	u, err := url.Parse(s.baseUrl)
	if err != nil {
		return SitronicsStationInfo{}, err
	}

	u.Path = path.Join(u.Path, "MainApi", "GetChargePointDescription")

	q := u.Query()
	q.Set("cPName", stationName)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return SitronicsStationInfo{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return SitronicsStationInfo{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SitronicsStationInfo{}, err
	}

	var result SitronicsStationInfo
	err = json.Unmarshal(body, &result)
	if err != nil {
		return SitronicsStationInfo{}, err
	}

	return result, nil

}
