package wordnik

import (
	"encoding/json"
	"net/http"
)

type WordOfTheDay struct {
	ContentProvider struct {
		ID   float64 `json:"id"`
		Name string  `json:"name"`
	} `json:"contentProvider"`
	Definitions []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Source       string `json:"source"`
		Text         string `json:"text"`
	} `json:"definitions"`
	Examples []struct {
		ID    float64 `json:"id"`
		Text  string  `json:"text"`
		Title string  `json:"title"`
		URL   string  `json:"url"`
	} `json:"examples"`
	ID          float64 `json:"id"`
	Note        string  `json:"note"`
	PublishDate string  `json:"publishDate"`
	Word        string  `json:"word"`
}

func unmarshalResponse(resp *http.Response, dst interface{}) error {
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(dst)
}
