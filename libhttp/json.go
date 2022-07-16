package libhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendJSON(ctx context.Context, w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Errorf(err.Error())
	}
}

func SendFile(ctx context.Context, w http.ResponseWriter, data []byte, filename, contentType string) {
	w.Header().Set("Content-type", contentType)
	w.Header().Set("Content-disposition", "attachment; filename="+filename)
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(data)
	if err != nil {
		fmt.Errorf(err.Error())
	}
}

// ReceiveJSON context будет использоваться в будущем при расширении функции.
func ReceiveJSON(ctx context.Context, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}
