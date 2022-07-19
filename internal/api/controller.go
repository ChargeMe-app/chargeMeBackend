package api

import (
	"fmt"
	"net/http"

	"github.com/poorfrombabylon/chargeMeBackend/specs/schema"
)

var _ schema.ServerInterface = &apiServer{}

type apiServer struct{}

func NewApiServer() schema.ServerInterface {
	return &apiServer{}
}

// Проверка сервиса
// (GET /healthz)
func (api *apiServer) HealthCheck(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	w.Write([]byte("hello healthCheck"))
	fmt.Println("hello healthCheck")
}
