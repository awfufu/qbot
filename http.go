// qbot/qbot.go
package qbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type cqResponse struct {
	Status  string          `json:"status"`
	Retcode int             `json:"retcode"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
	Wording string          `json:"wording"`
}

// Send raw parameters to NapCat
func (b *Bot) SendParams(action string, params map[string]any) (json.RawMessage, error) {
	resp, err := b.sendHttpRequest(action, params)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (b *Bot) sendHttpRequest(action string, params map[string]any) (*cqResponse, error) {
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("marshal params: %v", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, b.apiEndpoint+"/"+action, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, fmt.Errorf("create request: %v", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := b.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %v", err)
	}

	var cqResp cqResponse
	if err := json.Unmarshal(body, &cqResp); err != nil {
		return nil, fmt.Errorf("unmarshal response: %v", err)
	}

	return &cqResp, nil
}

func (b *Bot) handleHttpEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	jsonMap := make(map[string]any)
	if err := json.Unmarshal(body, &jsonMap); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if postType, exists := jsonMap["post_type"]; exists {
		if str, ok := postType.(string); ok && str != "" {
			go b.handleEvents(&str, &body, &jsonMap)
		}
	}
	w.WriteHeader(http.StatusOK)
}
