// qbot/qbot.go
package qbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type eventHeader struct {
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
}

type httpResponse struct {
	Status  string          `json:"status"`
	Retcode int             `json:"retcode"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
	Wording string          `json:"wording"`
}

// Send raw parameters to NapCat
func (s *Sender) SendParams(action string, params map[string]any) (json.RawMessage, error) {
	resp, err := s.sendHttpRequest(action, params)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (s *Sender) sendHttpRequest(action string, params map[string]any) (*httpResponse, error) {
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("marshal params: %v", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, s.apiEndpoint+"/"+action, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, fmt.Errorf("create request: %v", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := s.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %v", err)
	}

	var result httpResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response: %v", err)
	}

	return &result, nil
}

func (r *Receiver) handleHttpEvent(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	var header eventHeader
	if err := json.Unmarshal(body, &header); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if header.PostType != "" {
		go r.handleEvents(&header, &body)
	}
	w.WriteHeader(http.StatusOK)
}
