package wordnik

import (
	"fmt"
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

// RandomDictionaryWord returns a random dictionary word of any corpus count, length, and dictionary count.
func (c *APIClient) RandomDictionaryWord() (*Word, error) {
	return c.randomWord(true, 0, -1, 0, -1, 0, -1)
}

// RandomWord returns a random word of any corpus count, length, and dictionary count.
func (c *APIClient) RandomWord() (*Word, error) {
	return c.randomWord(false, 0, -1, 0, -1, 0, -1)
}

// RandomDictionaryWord returns a random dictionary word of any corpus count, specified min and max length, and any dictionary count.
func (c *APIClient) RandomDictionaryWordOfLength(minLength, maxLength int) (*Word, error) {
	return c.randomWord(true, 0, -1, minLength, maxLength, 0, -1)
}

// RandomWord returns a random word of any corpus count, specified min and max length, and any dictionary count.
func (c *APIClient) RandomWordOfLength(minLength, maxLength int) (*Word, error) {
	return c.randomWord(true, 0, -1, minLength, maxLength, 0, -1)
}

func (c *APIClient) randomWord(hasDictionaryDef bool,
	minCorpusCount, maxCorpusCount int,
	minLength, maxLength int,
	minDictionaryCount, maxDictionaryCount int) (*Word, error) {

	url := fmt.Sprintf("%s/randomWord?hasDictionaryDef=%t&minCorpusCount=%d&maxCorpusCount=%d&minLength=%d&maxLength=%d&minDictionaryCount=%d&maxDictionaryCount=%d",
		WORDS_BASE,
		minCorpusCount, maxCorpusCount,
		minLength, maxLength,
		minDictionaryCount, maxDictionaryCount)

	log.Printf(url)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(r)
	if err != nil {
		return nil, err
	}

	var word Word
	err = unmarshalResponse(resp, &word)

	return &word, err
}
