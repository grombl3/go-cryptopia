package cryptopia

import (
	"encoding/json"
)

type jsonResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"data"`
	Error string          `json:"Error"`
}
