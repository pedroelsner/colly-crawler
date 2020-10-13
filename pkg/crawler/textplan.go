package crawler

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
)

func GET(url string, hanlder HandleFunc) {

	// Delay request
	rand.Seed(time.Now().UnixNano())
	delay := rand.Intn(4) + 1
	fmt.Printf("Sleeping %d seconds...\n", delay)
	time.Sleep(time.Duration(delay) * time.Second)

	// Logger
	logger := logrus.WithField("url", url)
	start := time.Now()

	// Client
	c := colly.NewCollector()
	_ = c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 1 * time.Second,
	})

	// Handler
	c.OnHTML("html", func(e *colly.HTMLElement) {
		end := time.Now()
		duration := end.Sub(start)

		logger.
			WithField("duration", duration).
			Info("Requested")

		if err := hanlder(logger, e.DOM); err != nil {
			logger.Error(err)
		}
	})

	// Execution
	if err := c.Visit(url); err != nil {
		logger.Warn(err)
	}
}
