package qbot

import (
	"encoding/json"
	"net/http"
)

type Bot struct {
	httpClient    *http.Client
	httpServer    *http.Server
	apiEndpoint   string
	enableDebug   bool
	eventHandlers struct {
		groupMsg   []func(b *Bot, msg *Message)
		privateMsg []func(b *Bot, msg *Message)
	}
}

type cqResponse struct {
	Status  string          `json:"status"`
	Retcode int             `json:"retcode"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
	Wording string          `json:"wording"`
}
