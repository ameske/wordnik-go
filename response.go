package wordnik

import (
	"encoding/json"
	"net/http"
)

func unmarshalResponse(resp *http.Response, dst interface{}) error {
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(dst)
}
