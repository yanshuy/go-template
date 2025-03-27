package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func CreateStack(ms ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(ms) - 1; i >= 0; i-- {
			m := ms[i]
			next = m(next)
		}
		return next
	}
}

type responseWriterWithStatus struct {
	http.ResponseWriter
	statusCode int
}

func (rws *responseWriterWithStatus) WriteHeader(code int) {
	rws.ResponseWriter.WriteHeader(code)
	rws.statusCode = code
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rws := &responseWriterWithStatus{
			ResponseWriter: w,
		}
		next.ServeHTTP(rws, r)
		log.Println(r.Method, r.URL.Path, rws.statusCode, time.Since(start))
	})
}

func HiMom(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hi Mom!")
		next.ServeHTTP(w, r)
	})
}

func HiDad(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hi Dad!")
		next.ServeHTTP(w, r)
	})
}
