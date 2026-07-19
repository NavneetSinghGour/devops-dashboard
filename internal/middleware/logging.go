package middleware

import (
	"net/http"
	"time"

	"github.com/NavneetSinghGour/devops-dashboard/internal/logger"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(rw, r)

		logger.Info(
			r.Method +
				" " +
				r.URL.Path +
				" Status=" +
				http.StatusText(rw.statusCode) +
				" Duration=" +
				time.Since(start).String(),
		)
	})
}
