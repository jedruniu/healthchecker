package healthchecker

import (
	"fmt"
	"net/http"
)

type Server struct {
	Healths []RunReporter
}

func (s Server) HealthEndpoint(w http.ResponseWriter, r *http.Request) {
	var content string
	for _, health := range s.Healths {
		singleHealth := fmt.Sprintln(health, health.IsHealthy())
		content += singleHealth
	}
	w.Write([]byte(content))
}
