package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/aronkst/go-web-scraping/data"
	"github.com/aronkst/go-web-scraping/find"
	"github.com/julienschmidt/httprouter"
)

func HandlerHTML(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	body, err := getBody(request)
	if err != nil {
		httpError(writer, err)

		return
	}

	html, err := find.GetHTML(body)
	if err != nil {
		httpError(writer, err)

		return
	}

	output := data.OutputHTML{HTML: html}
	outputBody, err := json.Marshal(output)

	if err != nil {
		httpError(writer, err)

		return
	}

	_, err = writer.Write(outputBody)
	if err != nil {
		httpError(writer, err)

		return
	}
}

func HandlerFind(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	body, err := getBody(request)
	if err != nil {
		httpError(writer, err)

		return
	}

	var allFind []data.Find
	var allFindList []data.FindList

	for _, bodyFind := range body.Find {
		_, attributeOk := bodyFind["attribute"].(string)
		if attributeOk {
			find := data.Find{
				Name:      bodyFind["name"].(string),
				Class:     bodyFind["class"].(string),
				Attribute: bodyFind["attribute"].(string),
			}

			allFind = append(allFind, find)
		} else {
			var subFind []data.Find

			for _, sFind := range bodyFind["find"].([]interface{}) {
				switch sf := sFind.(type) {
				case map[string]any:
					find := data.Find{
						Name:      sf["name"].(string),
						Class:     sf["class"].(string),
						Attribute: sf["attribute"].(string),
					}

					subFind = append(subFind, find)
				}
			}

			findList := data.FindList{
				Name:  bodyFind["name"].(string),
				Class: bodyFind["class"].(string),
				Find:  subFind,
			}

			allFindList = append(allFindList, findList)
		}
	}

	var resultValue []any

	if len(allFind) >= 1 || len(allFindList) >= 1 {
		value, valueList, err := find.GetValues(body, allFind, allFindList)
		if err != nil {
			httpError(writer, err)

			return
		}

		for _, v := range value {
			resultValue = append(resultValue, v)
		}

		for _, vl := range valueList {
			resultValue = append(resultValue, vl)
		}
	}

	output := data.OutputFind{Find: resultValue}
	outputBody, err := json.Marshal(output)

	if err != nil {
		httpError(writer, err)

		return
	}

	_, err = writer.Write(outputBody)
	if err != nil {
		httpError(writer, err)

		return
	}
}

func getBody(request *http.Request) (data.Body, error) {
	rBody, err := io.ReadAll(request.Body)
	if err != nil {
		return data.Body{}, err
	}

	body := data.Body{}
	err = json.Unmarshal(rBody, &body)

	if err != nil {
		return data.Body{}, err
	}

	return body, nil
}

func httpError(writer http.ResponseWriter, err error) {
	http.Error(writer, err.Error(), http.StatusInternalServerError)
}
