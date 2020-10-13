package zukerman

import (
	"encoding/json"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"gorm.io/datatypes"

	"github.com/pedroelsner/colly-crawler/internal/app/format"
	"github.com/pedroelsner/colly-crawler/internal/repository/asset"
	"github.com/pedroelsner/colly-crawler/pkg/crawler"
)

func Detail(url string) {
	crawler.GETRendered(url, func(logger *logrus.Entry, dom *goquery.Selection) error {
		repository := asset.InitSqliteRepository()

		// Find by URL
		record, _ := repository.FindByURL(url)

		record.Source = "zukerman"
		record.URL = url

		// City, Zone, Neighborhood
		breadcrumb := dom.Find("ol.breadcrumb > li[itemprop] > a > span")

		if breadcrumb.Length() == 7 {
			// Example 01:
			// 0 Zukerman Leilões
			// 1 Imóveis
			// 2 SP
			// 3 São Paulo
			// 4 Zona Sul
			// 5 Vila Caraguatá
			// 6 Lote 001
			//
			// Example 02:
			// 0 Zukerman Leilões
			// 1 Imóveis
			// 2 SP
			// 3 Interior
			// 4 Sorocaba
			// 5 Jardim Santa Rosália
			// 6 Lote 001

			zone := ""

			breadcrumb.Each(func(i int, selection *goquery.Selection) {
				value := format.Trim(selection.Text())

				switch i {
				case 2:
					record.State = value
				case 3:
					if selection.Text() == "Interior" {
						zone = value
					} else {
						record.City = value
					}
				case 4:
					if zone != "" {
						record.City = value
					} else {
						zone = value
					}
				case 5:
					record.Neighborhood = value
				}
			})
		} else {

			// Example:
			// 0 Zukerman Leilões
			// 1 Imóveis
			// 2 AC
			// 3 Rio Branco
			// 4 Dom Giocondo
			// 5 Lote 018
			breadcrumb.Each(func(i int, selection *goquery.Selection) {
				value := format.Trim(selection.Text())

				switch i {
				case 2:
					record.State = value
				case 3:
					record.City = value
				case 4:
					record.Neighborhood = value
				}
			})
		}

		// Address
		record.Address = format.Trim(dom.Find("div.s-d-ld-i2").Contents().Not("span").Text())

		// Description
		record.Description = format.Trim(dom.Find("div.s-d-ld-i1").Text())

		// Tags
		tags := make(map[string]string)
		dom.Find("div.s-d-ld-i3 p").Each(func(i int, selection *goquery.Selection) {
			key := strings.Trim(selection.Find("span").Text(), ":")
			value := format.Trim(selection.Contents().Not("span").Text())

			// Intercept type
			if key == "Tipo" {
				record.Type = value
				return
			}

			// Link Processo
			if key == "Processo" {
				link, _ := selection.Find("a[href]").First().Attr("href")
				tags["Link Processo"] = link
			}

			tags[key] = value
		})

		tagsJson, _ := json.Marshal(tags)
		record.Tags = datatypes.JSON(tagsJson)

		// Documents
		documents := make(map[string]string)
		dom.Find("div.s-d-do-i-main p").Each(func(i int, selection *goquery.Selection) {
			key := format.Trim(selection.Find("a[href]").First().Contents().Not("span").Text())
			value, _ := selection.Find("a[href]").First().Attr("href")

			documents[key] = value
		})

		documentsJson, _ := json.Marshal(documents)
		record.Documents = datatypes.JSON(documentsJson)

		// Owner
		record.Owner = format.Trim(dom.Find("div.d-n-v > h2").Text())

		// LastBid
		record.LastBid = format.Currency(dom.Find("div.s-d-lb3 > div.m-l-o").Text())

		// IncrementBid
		record.IncrementBid = format.Currency(
			strings.ReplaceAll(dom.Find("div.s-d-lb6-2").Text(), "Incremento mínimo", ""),
		)

		// First
		first := dom.Find("div.s-d-lb5-1")
		record.FirstDate = format.Trim(first.Find("div.daet").Text())
		record.FirstPrice = format.Currency(first.Find("span.dvla").Contents().Not("div").Text())

		// Second
		second := dom.Find("div.s-d-lb5-2")
		record.SecondDate = format.Trim(second.Find("div.daet").Text())
		record.SecondPrice = format.Currency(second.Find("span.dvld").Contents().Not("div").Text())

		// Save
		return repository.Save(record)
	})
}
