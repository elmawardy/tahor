package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func Log(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := randomHex(5)

		log.WithFields(log.Fields{
			"Source IP ": r.RemoteAddr,
			"Path":       r.URL,
			"Id":         id,
		}).Info(r.Method)

		start := time.Now()
		next.ServeHTTP(w, r)
		end := time.Now()

		log.WithFields(log.Fields{
			"Source IP ": r.RemoteAddr,
			"Path":       r.URL,
			"Time":       end.Sub(start),
			"Id":         id,
		}).Info(r.Method)
	}
}
