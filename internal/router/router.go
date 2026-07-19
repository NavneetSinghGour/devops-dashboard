package router

import (
	"net/http"

	"github.com/NavneetSinghGour/devops-dashboard/internal/handlers"
	"github.com/NavneetSinghGour/devops-dashboard/internal/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRoutes() {

	mux := http.NewServeMux()

	// Static Files
	fs := http.FileServer(http.Dir("internal/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Dashboard
	mux.HandleFunc("/", handlers.Dashboard)

	// Health Endpoints
	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/ready", handlers.Health)
	mux.HandleFunc("/live", handlers.Health)

	// Prometheus Metrics
	mux.Handle("/metrics", promhttp.Handler())

	// APIs
	mux.HandleFunc("/api/info", handlers.ApplicationInfo)
	mux.HandleFunc("/api/build", handlers.BuildInfo)
	mux.HandleFunc("/api/runtime", handlers.RuntimeInfo)
	mux.HandleFunc("/api/kubernetes", handlers.KubernetesInfo)

	// Middleware
	http.Handle("/", middleware.Logging(mux))
}
