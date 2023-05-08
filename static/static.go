package static

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetHTML(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		errorMessage := fmt.Sprintf("status code error %d, %s", response.StatusCode, response.Status)
		return "", errors.New(errorMessage)
	}

	html, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(html), nil
}
