package qbot

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Bot struct {
	httpClient    *http.Client
	httpServer    *http.Server
	apiEndpoint   string
	eventHandlers struct {
		message   []func(b *Bot, msg *Message)
		emojiLike []func(b *Bot, msg *EmojiReaction)
		recall    []func(b *Bot, msg *RecallNotice)
		poke      []func(b *Bot, msg *PokeNotify)
	}
}

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
	bot.eventHandlers.message = make([]func(b *Bot, msg *Message), 0)
	bot.eventHandlers.emojiLike = make([]func(b *Bot, msg *EmojiReaction), 0)
	bot.eventHandlers.recall = make([]func(b *Bot, msg *RecallNotice), 0)
	bot.eventHandlers.poke = make([]func(b *Bot, msg *PokeNotify), 0)

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

func (b *Bot) OnMessage(handler func(b *Bot, msg *Message)) {
	b.eventHandlers.message = append(b.eventHandlers.message, handler)
}

func (b *Bot) OnEmojiReaction(handler func(b *Bot, emoji *EmojiReaction)) {
	b.eventHandlers.emojiLike = append(b.eventHandlers.emojiLike, handler)
}

func (b *Bot) OnRecall(handler func(b *Bot, recall *RecallNotice)) {
	b.eventHandlers.recall = append(b.eventHandlers.recall, handler)
}

func (b *Bot) OnPoke(handler func(b *Bot, poke *PokeNotify)) {
	b.eventHandlers.poke = append(b.eventHandlers.poke, handler)
}
