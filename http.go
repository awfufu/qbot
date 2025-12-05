// qbot/qbot.go
package qbot

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

// Send raw parameters to NapCat
func (b *Bot) SendParams(action string, params map[string]any) (json.RawMessage, error) {
	if b.enableDebug {
		log.Println()
		jsonBytes, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		log.Printf("[Debug] qbot.SendParams: %s\n%s", action, string(jsonBytes))
	}
	resp, err := b.sendHttpRequest(action, params)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (b *Bot) sendHttpRequest(action string, params map[string]any) (*cqResponse, error) {
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	var httpResp *http.Response
	var reqErr error

	// Retry logic from sendRequest
	for i := range 3 {
		httpReq, err := http.NewRequest(http.MethodPost, b.apiEndpoint+"/"+action, bytes.NewBuffer(jsonBytes))
		if err != nil {
			return nil, err
		}
		httpReq.Header.Set("Content-Type", "application/json")

		httpResp, reqErr = b.httpClient.Do(httpReq)
		if reqErr == nil {
			break // Successfully sent, exit retry loop
		}
		log.Printf("Request failed: %v. Retrying (%d/3)...", reqErr, i+1)
		time.Sleep(1 * time.Second)
	}

	if reqErr != nil { // If all retries failed
		log.Printf("Request failed: %v", reqErr)
		return nil, reqErr
	}

	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		log.Printf("Request failed: %v", err)
		return nil, err
	}

	if b.enableDebug {
		log.Printf("[Debug] %s: %s\n%s", action, httpResp.Status, string(body))
	}

	var cqResp cqResponse
	if err := json.Unmarshal(body, &cqResp); err != nil {
		return nil, err
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
