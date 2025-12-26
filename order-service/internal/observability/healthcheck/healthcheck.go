package health

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthStatus struct {
	Status       string            `json:"status"`
	Service      string            `json:"service"`
	Timestamp    time.Time         `json:"timestamp"`
	Dependencies map[string]string `json:"dependencies,omitempty"`
}

func NewHealthHandler(serviceName string, dependencies map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := HealthStatus{
			Status:       "ok",
			Service:      serviceName,
			Timestamp:    time.Now().UTC(),
			Dependencies: dependencies,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(status)
	}
}
