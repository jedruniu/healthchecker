package healthcheck

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	Healths []HealthChecker
}

func (s Server) HealthEndpoint(w http.ResponseWriter, r *http.Request) {
	report := make(map[string]bool)
	for _, health := range s.Healths {
		report[fmt.Sprint(health)] = health.IsHealthy()
	}
	// not being able to marshal map seems impossible, skipping error
	content, _ := json.Marshal(report)
	w.Write(content)
}
