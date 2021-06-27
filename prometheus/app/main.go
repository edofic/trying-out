package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	addr := ":2112"
	log.Println("Starting server on", addr)
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", newHandleReq())
	http.ListenAndServe(addr, nil)
}

type handleReq struct {
	reqCounter prometheus.Counter
}

func newHandleReq() http.Handler {
	return &handleReq{
		reqCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "app_req_counter",
			Help: "The total number of processed requests",
		}),
	}
}

func (h *handleReq) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.reqCounter.Inc()
	w.Write([]byte("ok\n"))
}
