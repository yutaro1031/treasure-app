package httputil

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/pkg/errors"

	"github.com/yutaro1031/treasure-app/mid-task/backend/model"
)

func Search(keyword string) (interface{}, error) {
	rakutenAPIKey := os.Getenv("RAKUTEN_API_KEY")
	rakutenURL := os.Getenv("RAKUTEN_URL")

	values := url.Values{}
	values.Add("applicationId", rakutenAPIKey)

	resp, _ := http.Get(rakutenURL + "?" + values.Encode())
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	jsonBytes := ([]byte)(byteArray)
	data := new(model.SearchedBooks)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		return nil, errors.Wrap(err, "JSON Unmarshal error")
	}

	return data, nil
}
