package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

type HandleFunc func(*logrus.Entry, *goquery.Selection) error
