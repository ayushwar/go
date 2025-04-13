package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody reads the request body and decodes the JSON into the provided struct
func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return // stop if reading body fails
	}
	
	err = json.Unmarshal(body, x)
	if err != nil {
		return // stop if unmarshalling fails
	}
}
