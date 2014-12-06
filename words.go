package wordnik

import (
	"log"
	"net/http"
	"time"
)

const (
	WORDS_BASE = BASE + "/words.json/"
)

// WordOfTheDay fetches the Wordnik wotd for the given day. It returns a struct representing
// the JSON response documented by the Wodnik developer guide.
func (c *APIClient) WordOfTheDay(date time.Time) (*WordOfTheDay, error) {
	url := WORDS_BASE + "wordOfTheDay?date=" + date.Format("2006-01-02")

	log.Printf(url)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(r)
	if err != nil {
		return nil, err
	}

	var wotd WordOfTheDay
	err = unmarshalResponse(resp, &wotd)

	return &wotd, err
}
