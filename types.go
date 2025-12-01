package qbot

import (
	"encoding/json"
	"net/http"
)

type Bot struct {
	httpClient    *http.Client
	httpServer    *http.Server
	apiEndpoint   string
	eventHandlers struct {
		groupMsg   []func(bot *Bot, msg *Message)
		privateMsg []func(bot *Bot, msg *Message)
	}
}

type cqRequest struct {
	Action string         `json:"action"`
	Params map[string]any `json:"params"`
}

type cqResponse struct {
	Status  string          `json:"status"`
	Retcode int             `json:"retcode"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
	Wording string          `json:"wording"`
}
