package wordnik

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

var testWOTD = `{"contentProvider":{"id":711,"name":"wordnik"},"definitions":[{"partOfSpeech":"noun","source":"century","text":"Formerly, in Maine and some other parts of New England, a house-cleaning party; a gathering of neighbors to aid one of their number in cleaning house."}],"examples":[{"id":1.148981134e+09,"text":"Even those evil days of New England households, the annual house-cleaning, were robbed of some of their dismal terrors by what was known as a \"whang,\" a gathering of a few friendly women neighbors to assist one another in that dire time, and thus speed and shorten the hours of misery.","title":"Home Life in Colonial Days","url":"http://api.wordnik.com/v4/mid/f3ce9fa9c63071b4d482f2e8407175ae9000341f82964714f5756e154741cb3e"}],"id":520329,"note":"The word 'whang' is related to the word 'thong'.","publishDate":"2014-12-06T03:00:00.000+0000","word":"whang"}`

var (
	testKey string
)

// In order to run tests we have to have an api key. Read it in from wordnik.json
func init() {
	bytes, err := ioutil.ReadFile("wordnik.json")
	if err != nil {
		log.Fatalf(err.Error())
	}

	config := struct {
		ApiKey string
	}{}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	testKey = config.ApiKey
}

func TestWordOfTheDay(t *testing.T) {
	c := NewAPIClient(testKey)

	wotd, err := c.WordOfTheDay(time.Now())
	if err != nil {
		t.Error(err.Error())
	}

	bytes, err := json.Marshal(wotd)
	if err != nil {
		t.Error(err.Error())
	}

	if string(bytes) != testWOTD {
		t.Errorf("TestWordOfTheDay returned unexpected response:\n\t%s", string(bytes))
	}
}
