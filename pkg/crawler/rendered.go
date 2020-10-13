package crawler

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/sirupsen/logrus"
)

func GETRendered(url string, handler HandleFunc) {

	// Delay request
	rand.Seed(time.Now().UnixNano())
	delay := rand.Intn(4) + 1
	fmt.Printf("Sleeping %d seconds...\n", delay)
	time.Sleep(time.Duration(delay) * time.Second)

	// Logger
	logger := logrus.WithField("url", url)
	start := time.Now()

	// Client
	geziyor.NewGeziyor(&geziyor.Options{
		LogDisabled: true,
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered(url, g.Opt.ParseFunc)
		},

		// Handler
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			end := time.Now()
			duration := end.Sub(start)

			logger.
				WithField("duration", duration).
				Info("Requested")

			if err := handler(logger, r.HTMLDoc.Find("html")); err != nil {
				logger.Error(err)
			}
		},
	}).Start()
}
