// qbot/qbot.go
package qbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func NewBot(address string) *Bot {
	bot := &Bot{
		httpClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:        10,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
				DisableKeepAlives:   false,
			},
			Timeout: 10 * time.Second,
		},
	}
	bot.eventHandlers.groupMsg = make([]func(b *Bot, msg *Message), 0)
	bot.eventHandlers.privateMsg = make([]func(b *Bot, msg *Message), 0)

	bot.httpServer = &http.Server{
		Addr:         address,
		Handler:      http.HandlerFunc(bot.handleHttpEvent),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return bot
}

func (b *Bot) ConnectNapcat(url string) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		log.Fatal("Invalid URL")
	}
	url = strings.TrimRight(url, "/")
	b.apiEndpoint = url

	// Initial handshake with retry
	for {
		resp, err := b.httpClient.Get(url)
		if err != nil {
			log.Printf("Failed to connect to NapCat: %v. Retrying in 3 seconds...", err)
			time.Sleep(3 * time.Second)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response from NapCat: %v. Retrying in 3 seconds...", err)
			time.Sleep(3 * time.Second)
			continue
		}

		var cqResp cqResponse
		if err := json.Unmarshal(body, &cqResp); err != nil {
			log.Printf("Failed to parse response from NapCat: %v. Retrying in 3 seconds...", err)
			time.Sleep(3 * time.Second)
			continue
		}

		if cqResp.Status == "ok" && strings.Contains(cqResp.Message, "NapCat") {
			log.Printf("Connected to NapCat: %s", cqResp.Message)
			break
		} else {
			log.Printf("Unexpected response from NapCat: %s. Retrying in 3 seconds...", string(body))
			time.Sleep(3 * time.Second)
			continue
		}
	}
}

func (b *Bot) Run() error {
	log.Printf("Listening on %s", b.httpServer.Addr)
	if err := b.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (b *Bot) GroupMsg(handler func(b *Bot, msg *Message)) {
	b.eventHandlers.groupMsg = append(b.eventHandlers.groupMsg, handler)
}

func (b *Bot) PrivateMsg(handler func(b *Bot, msg *Message)) {
	b.eventHandlers.privateMsg = append(b.eventHandlers.privateMsg, handler)
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

func (b *Bot) sendRequest(req *cqRequest) (*http.Response, error) {
	jsonBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, err
	}

	var resp *http.Response
	var reqErr error

	// Retry logic
	for i := range 3 {
		httpReq, err := http.NewRequest(http.MethodPost, b.apiEndpoint+"/"+req.Action, bytes.NewBuffer(jsonBytes))
		if err != nil {
			return nil, err
		}
		httpReq.Header.Set("Content-Type", "application/json")

		resp, reqErr = b.httpClient.Do(httpReq)
		if reqErr == nil {
			return resp, nil
		}
		log.Printf("Request failed: %v. Retrying (%d/3)...", reqErr, i+1)
		time.Sleep(1 * time.Second)
	}

	return nil, reqErr
}

func (b *Bot) sendWithResponse(req *cqRequest) (*cqResponse, error) {
	resp, err := b.sendRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d %s", resp.StatusCode, string(body))
	}

	var cqResp cqResponse
	if err := json.Unmarshal(body, &cqResp); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return &cqResp, nil
}

// Send implements api.Client interface
func (b *Bot) Send(action string, params map[string]any) (json.RawMessage, error) {
	req := cqRequest{
		Action: action,
		Params: params,
	}
	resp, err := b.sendWithResponse(&req)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
