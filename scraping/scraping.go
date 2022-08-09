package scraping

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/aronkst/go-web-scraping/data"
)

func GetValues(html string, find []data.Find, findList []data.FindList) ([]data.Value, []data.ValueList, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, nil, err
	}

	var value []data.Value

	for _, f := range find {
		var v string

		doc.Find(f.Class).Each(func(_ int, s *goquery.Selection) {
			if f.Attribute == "" {
				v = s.Text()
			} else {
				v, _ = s.Attr(f.Attribute)
			}
		})

		value = append(value, data.Value{Name: f.Name, Value: v})
	}

	var valueList []data.ValueList

	for _, fl := range findList {
		var arraySubValue [][]data.Value

		doc.Find(fl.Class).Each(func(i int, s *goquery.Selection) {
			var subValue []data.Value

			for _, f := range fl.Find {
				var v string

				s.Find(f.Class).Each(func(_ int, sa *goquery.Selection) {
					if f.Attribute == "" {
						v = sa.Text()
					} else {
						v, _ = sa.Attr(f.Attribute)
					}

					subValue = append(subValue, data.Value{Name: f.Name, Value: v})
				})
			}

			arraySubValue = append(arraySubValue, subValue)
		})

		valueList = append(valueList, data.ValueList{Name: fl.Name, Values: arraySubValue})
	}

	return value, valueList, nil
}
