package qbot

import (
	"net/http"
	"strings"
	"time"
)

type Sender struct {
	httpClient  *http.Client
	apiEndpoint string
}

type Receiver struct {
	httpServer *http.Server

	// Channels for events
	message   chan *Message
	emojiLike chan *EmojiReaction
	recall    chan *RecallNotice
	poke      chan *PokeNotify
	err       chan error // Channel for server errors
}

func HttpClient(url string) *Sender {
	url = strings.TrimRight(url, "/")
	return &Sender{
		httpClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:        10,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
				DisableKeepAlives:   false,
			},
			Timeout: 10 * time.Second,
		},
		apiEndpoint: url,
	}
}

func HttpServer(address string) *Receiver {
	rx := &Receiver{
		message:   make(chan *Message, 100),
		emojiLike: make(chan *EmojiReaction, 100),
		recall:    make(chan *RecallNotice, 100),
		poke:      make(chan *PokeNotify, 100),
		err:       make(chan error, 1),
	}

	rx.httpServer = &http.Server{
		Addr:         address,
		Handler:      http.HandlerFunc(rx.handleHttpEvent),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start server in goroutine
	go func() {
		if err := rx.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			rx.err <- err
		}
		close(rx.err)
		close(rx.message)
		close(rx.emojiLike)
		close(rx.recall)
		close(rx.poke)
	}()

	return rx
}

// Event Channels

func (r *Receiver) OnMessage() <-chan *Message {
	return r.message
}

func (r *Receiver) OnEmojiReaction() <-chan *EmojiReaction {
	return r.emojiLike
}

func (r *Receiver) OnRecall() <-chan *RecallNotice {
	return r.recall
}

func (r *Receiver) OnPoke() <-chan *PokeNotify {
	return r.poke
}

func (r *Receiver) Error() <-chan error {
	return r.err
}

func (r *Receiver) Close() error {
	return r.httpServer.Close()
}
