package api

import "encoding/json"

type Client interface {
	SendParams(action string, params map[string]any) (json.RawMessage, error)
}
