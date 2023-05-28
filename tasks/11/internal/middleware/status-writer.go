package middleware

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

// StatusWriter is a custom ResponseWriter with capture status and size of response
type StatusWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (w *StatusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
func (w *StatusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.size += n
	return n, err
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := StatusWriter{ResponseWriter: w}

		next(&sw, r)

		fields := []zap.Field{
			zap.Int("status", sw.status),
			zap.Int("response-size", sw.size),
			zap.String("method", r.Method),
			zap.String("host", r.Host),
			zap.String("remote-ip", r.RemoteAddr),
			zap.String("uri", r.RequestURI),
			zap.String("latency", time.Since(start).String()),
		}

		n := sw.status
		switch {
		case n >= 500:
			zap.L().Error("[SERVER ERROR]", fields...)
		case n >= 400:
			zap.L().Warn("[CLIENT ERROR]", fields...)
		case n >= 300:
			zap.L().Info("[Redirect]", fields...)
		default:
			zap.L().Info("[Success]", fields...)
		}

	}
}
