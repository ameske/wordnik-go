package wordnik

import (
	"fmt"
	"log"
	"net/http"
)

const (
	WORD_BASE = BASE + "/word.json"
)

const (
	All        = "all"
	Ahd        = "ahd"
	Century    = "century"
	Wiktionary = "wiktionary"
	Webster    = "webster"
	Wordnet    = "wordnet"
)

func (c *APIClient) Definition(word string) ([]Definition, error) {
	return c.definition(word, 200, false, "all", false)
}

func (c *APIClient) DefinitionCanonical(word string) ([]Definition, error) {
	return c.definition(word, 200, false, "all", true)
}

func (c *APIClient) DefinitionByDictionary(word string, dictionary string) ([]Definition, error) {
	return c.definition(word, 200, false, dictionary, false)
}

func (c *APIClient) DefinitionByDictionaryCanonical(word string, dictionary string) ([]Definition, error) {
	return c.definition(word, 200, false, dictionary, true)

}

func (c *APIClient) definition(word string, limit int, includeRelated bool, sourceDictionaries string, useCanonical bool) ([]Definition, error) {
	url := fmt.Sprintf("%s/%s/definitions?limit=%d&includeRelated=%t&sourceDictionaries=%s&useCanonical=%t&includeTags=false", WORD_BASE, word, limit, includeRelated, sourceDictionaries, useCanonical)
	log.Printf(url)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(r)
	if err != nil {
		return nil, err
	}

	var d []Definition
	err = unmarshalResponse(resp, &d)

	return d, err
}
