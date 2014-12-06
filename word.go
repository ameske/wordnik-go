package wordnik

import (
	"fmt"
	"net/http"
)

const (
	WORD_BASE = BASE + "/word.json/"
)

func (c *APIClient) Definition(word string) (*Definition, error) {
	return c.definition(word, 200, false, "all", false)
}

func (c *APIClient) DefinitionCanonical(word string) (*Definition, error) {
	return c.definition(word, 200, true, "all", true)
}

func (c *APIClient) definition(word string, limit int, includeRelated bool, sourceDictionaries string, useCanonical bool) (*Definition, error) {
	url := fmt.Sprintf("%s/definitions?limit=%d&includeRelated=%t&sourceDictionaries=%s&useCanonical=%t&includeTags=false", word, limit, sourceDictionaries, useCanonical)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(r)
	if err != nil {
		return nil, err
	}

	var d Definition
	err = unmarshalResponse(resp, &d)

	return &d, err
}
