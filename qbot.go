package qbot

import (
	"encoding/json"
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
			log.Printf("Connect to NapCat: %v", err)
			time.Sleep(3 * time.Second)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Read response: %v", err)
			time.Sleep(3 * time.Second)
			continue
		}

		var cqResp cqResponse
		if err := json.Unmarshal(body, &cqResp); err != nil {
			log.Printf("Parse response: %v", err)
			time.Sleep(3 * time.Second)
			continue
		}

		if cqResp.Status == "ok" && strings.Contains(cqResp.Message, "NapCat") {
			log.Printf("Connected to NapCat: %s", cqResp.Message)
			break
		} else {
			log.Printf("Unexpected response: %s", string(body))
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

func (b *Bot) Debug(status bool) {
	b.enableDebug = status
}
