package common_utils

import (
	"fmt"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	"github.com/poorfrombabylon/chargeMeBackend/internal/integration/my_ecars"
	"github.com/poorfrombabylon/chargeMeBackend/internal/integration/sitronics"
	"strconv"
	"strings"
)

type Connector struct {
	OutletID  outletDomain.OutletID
	Available int
}

type Station struct {
	StationID  stationDomain.StationID
	Connectors []Connector
}

func ConvertSitronicsStationToCommonStation(station sitronics.SitronicsStation) Station {
	stationID := stationDomain.StationID(station.Id)
	connectors := make([]Connector, 0, len(station.Connectors))

	for i, c := range station.Connectors {
		available := 0

		if c.StatusInt == 1 {
			available = 1
		}

		conn := Connector{
			OutletID:  outletDomain.OutletID(stationID.String() + strconv.Itoa(i)),
			Available: available,
		}

		connectors = append(connectors, conn)
	}

	return Station{
		StationID:  stationID,
		Connectors: connectors,
	}
}

func ConvertMyECarsStationToCommonStation(station my_ecars.MyECarsStation) (Station, error) {
	stationID := stationDomain.StationID(station.Id)
	connectors := make([]Connector, 0, len(station.Connectors))

	for i, c := range station.Connectors {
		available := 0

		availableString := fmt.Sprintf("%v", c.State)
		availableString = availableString[1 : len(availableString)-1]
		availableFields := strings.Split(availableString, " ")

		available, err := strconv.Atoi(availableFields[0])
		if err != nil {
			return Station{}, err
		}

		if available == 46 {
			available = 1
		}

		conn := Connector{
			OutletID:  outletDomain.OutletID(stationID.String() + strconv.Itoa(i)),
			Available: available,
		}

		connectors = append(connectors, conn)
	}

	return Station{
		StationID:  stationID,
		Connectors: connectors,
	}, nil
}
