package zukerman

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"

	"github.com/pedroelsner/colly-crawler/pkg/crawler"
)

func List(url string) {
	crawler.GETRendered(url, func(logger *logrus.Entry, dom *goquery.Selection) error {

		// Items
		dom.Find("div.cd-0").Each(func(i int, selection *goquery.Selection) {
			link, _ := selection.Find("a[href]").First().Attr("href")
			Detail(link)
		})

		// Pagination
		// dom.Find("ul.pagination").First().Find("li > a[rel='next']").Each(func(i int, selection *goquery.Selection) {
		// 	link, _ := selection.Attr("href")
		// 	List(link)
		// })

		return nil
	})
}
