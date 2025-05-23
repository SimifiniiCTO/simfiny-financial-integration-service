package api

import (
	"net/http"

	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/version"
)

// Version godoc
// @Summary Version
// @Description returns fis service version and git commit hash
// @Tags HTTP API
// @Produce json
// @Router /version [get]
// @Success 200 {object} api.MapResponse
func (s *Server) versionHandler(w http.ResponseWriter, r *http.Request) {
	result := map[string]string{
		"version": version.VERSION,
		"commit":  version.REVISION,
	}
	s.JSONResponse(w, r, result)
}
