package find

import (
	"github.com/aronkst/go-web-scraping/data"
	"github.com/aronkst/go-web-scraping/javascript"
	"github.com/aronkst/go-web-scraping/scraping"
	"github.com/aronkst/go-web-scraping/static"
)

func GetHTML(body data.Body) (string, error) {
	var html string
	var err error

	if body.Javascript {
		html, err = javascript.GetHTML(body.Url)
	} else {
		html, err = static.GetHTML(body.Url)
	}

	if err != nil {
		return "", err
	}

	return html, nil
}

func GetValues(body data.Body, find []data.Find, findList []data.FindList) ([]data.Value, []data.ValueList, error) {
	var html string
	var err error

	if body.HTML == "" {
		html, err = GetHTML(body)
	} else {
		html = body.HTML
	}

	if err != nil {
		return nil, nil, err
	}

	return scraping.GetValues(html, find, findList)
}
