package http

import (
	"net/http"

	"github.com/elena9bulat-wq/go-rest-userprofile/internal/observability"
	"github.com/elena9bulat-wq/go-rest-userprofile/internal/service"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	observability.InitMetrics()

	profileSvc := service.NewProfileService()
	h := NewHandler(profileSvc)

	// endpoint metrics
	mux.Handle("/metrics", observability.MetricsHandler())

	// endpoint instrumentat
	mux.HandleFunc("/profile", observability.InstrumentHandler("/profile", h.GetProfile))

	return mux
}