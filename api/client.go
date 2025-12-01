package api

import "encoding/json"

type Client interface {
	Send(action string, params map[string]any) (json.RawMessage, error)
}
