package router

import (
	"log"
	"net/http"
	"time"
)

// LoggingResponseWriter will encapsulate a standard ResponseWritter with a copy of its statusCode
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// ResponseWriterWrapper is supposed to capture statusCode from ResponseWriter
func ResponseWriterWrapper(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

// WriteHeader is a surcharge of the ResponseWriter method
func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// Logger is a gorilla/mux middleware to add log to the API
func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := ResponseWriterWrapper(w)
		inner.ServeHTTP(wrapper, r)
		// 127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326 286.219µs
		log.Printf("%s %s %s [%v] \"%s %s %s\" %d %d \"%s\" %s",
			r.RemoteAddr,
			"-",
			"-",
			start,
			r.Method,
			r.RequestURI,
			r.Proto, // string "HTTP/1.1"
			wrapper.statusCode,
			r.ContentLength,
			r.Header["User-Agent"],
			time.Since(start),
		)
	})
}
