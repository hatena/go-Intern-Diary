package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/hatena/go-Intern-Diary/model"
)

func GetCategoriesApi(tag_names []string) ([]*model.TagWithCategoriesJson, error) {
	values := buildParameter(tag_names)
	flaskUrl := "http://host.docker.internal:5000/categorize"
	rsp, err := http.Get(flaskUrl + "?" + values.Encode())
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	var categoryIDResponse []*model.TagWithCategoriesJson
	err = json.Unmarshal(body, &categoryIDResponse)
	if err != nil {
		return nil, err
	}
	return categoryIDResponse, nil
}

func buildParameter(tag_names []string) url.Values {
	values := url.Values{}
	for _, tag := range tag_names {
		values.Add("tag_name", tag)
	}
	return values
}
