package api

import (
	"chargeMe/specs/schema"
)

var _ schema.ServerInterface = &apiServer{}

type apiServer struct{}

func NewApiServer() schema.ServerInterface {
	return &apiServer{}
}
