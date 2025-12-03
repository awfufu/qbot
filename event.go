package qbot

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/awfufu/qbot/api"
)

func (b *Bot) handleEvents(postType *string, msgStr *[]byte, jsonMap *map[string]any) {
	switch *postType {
	case "meta_event":
		// heartbeat, connection state..
	case "notice":
		// TODO
		switch (*jsonMap)["notice_type"] {
		case "group_recall":
			// TODO
		}
	case "message":
		switch (*jsonMap)["message_type"] {
		case "private":
			msgJson := &api.MessageJson{}
			if json.Unmarshal(*msgStr, msgJson) != nil {
				return
			}
			if msg := parseMsgJson(msgJson); msg != nil {
				for _, handler := range b.eventHandlers.privateMsg {
					handler(b, msg)
				}
			}
		case "group":
			msgJson := &api.MessageJson{}
			if json.Unmarshal(*msgStr, msgJson) != nil {
				return
			}
			if msg := parseMsgJson(msgJson); msg != nil {
				for _, handler := range b.eventHandlers.groupMsg {
					handler(b, msg)
				}
			}
		}
	}
}

func parseMsgJson(raw *api.MessageJson) *Message {
	if raw == nil {
		return nil
	}
	result := Message{
		GroupID:  raw.GroupID,
		UserID:   raw.Sender.UserID,
		ReplyID:  0,
		Nickname: raw.Sender.Nickname,
		Card:     raw.Sender.Card,
		Role:     raw.Sender.Role,
		Time:     raw.Time,
		MsgID:    raw.MessageID,
		Raw:      raw.RawMessage,
	}
	for _, msg := range raw.Message {
		var jsonData map[string]any
		if err := json.Unmarshal(msg.Data, &jsonData); err != nil {
			return nil
		}
		switch msg.Type {
		case "text":
			if text, ok := jsonData["text"].(string); ok {
				result.Array = append(result.Array, &TextItem{
					Content: text,
				})
			}
		case "at":
			var qqStr string
			if qq, ok := jsonData["qq"].(string); ok {
				qqStr = qq
			} else if qq, ok := jsonData["qq"].(float64); ok {
				qqStr = fmt.Sprintf("%.0f", qq)
			}
			if qqStr != "" {
				qqInt, err := strconv.ParseUint(qqStr, 10, 64)
				if err != nil {
					continue
				}
				result.Array = append(result.Array, &AtItem{
					TargetID: qqInt,
				})
			}
		case "face":
			var idStr string
			if id, ok := jsonData["id"].(string); ok {
				idStr = id
			} else if id, ok := jsonData["id"].(float64); ok {
				idStr = fmt.Sprintf("%.0f", id)
			}
			if idStr != "" {
				idInt, err := strconv.ParseUint(idStr, 10, 64)
				if err != nil {
					continue
				}
				result.Array = append(result.Array, &FaceItem{
					ID: idInt,
				})
			}
		case "image":
			if url, ok := jsonData["url"].(string); ok {
				result.Array = append(result.Array, &ImageItem{
					URL: url,
				})
			}
		case "record":
			if path, ok := jsonData["path"].(string); ok {
				result.Array = append(result.Array, &RecordItem{
					Path: path,
				})
			}
		case "reply":
			if id, ok := jsonData["id"].(string); ok {
				result.ReplyID, _ = strconv.ParseUint(id, 10, 64)
			} else if id, ok := jsonData["id"].(float64); ok {
				result.ReplyID = uint64(id)
			}
		case "file":
			result.Array = append(result.Array, &FileItem{
				Data: string(msg.Data),
			})
		case "forward":
			result.Array = append(result.Array, &ForwardItem{
				Data: string(msg.Data),
			})
		case "json":
			result.Array = append(result.Array, &JsonItem{
				Data: string(msg.Data),
			})
		default:
			result.Array = append(result.Array, &OtherItem{
				Data: string(msg.Data),
			})
		}
	}
	return &result
}
