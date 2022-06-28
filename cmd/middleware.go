package main

import (
	"golambda/internal/logger"
	"net/http"
	"runtime/debug"
	"time"
)

func APIGatewayMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = r.URL.Path[4:]
		next.ServeHTTP(w, r)
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			_, err := w.Write([]byte{})
			if err != nil {
				logger.Error(err.Error())
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				logger.Infow("",
					"err", err,
					"trace", debug.Stack(),
				)
			}
		}()

		start := time.Now()
		rw := &ResponseWriter{ResponseWriter: w}
		next.ServeHTTP(rw, r)
		logger.Infow(
			"request summary",
			"status", rw.StatusCode,
			"method", r.Method,
			"path", r.URL.EscapedPath(),
			"duration", time.Since(start),
		)
	})
}
