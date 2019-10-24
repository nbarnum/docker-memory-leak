package main

import (
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	memoryTicker := time.NewTicker(time.Millisecond * 5)
	leak := make(map[int][]byte)
	i := 0

	go func() {
		for range memoryTicker.C {
			leak[i] = make([]byte, 1024)
			i++
		}
	}()

	fileTicker := time.NewTicker(time.Millisecond * 5)
	go func() {
		f, _ := os.Create("/tmp/file")
		buffer := make([]byte, 1024)
		defer f.Close()

		for range fileTicker.C {
			f.Write(buffer)
			f.Sync()
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8081", nil)
}
