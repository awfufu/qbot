package qbot

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/awfufu/qbot/api"
)

func (r *Receiver) handleEvents(header *eventHeader, msgStr *[]byte) {
	log.Println(string(*msgStr))
	switch header.PostType {
	case "notice":
		switch header.NoticeType {
		case "group_msg_emoji_like":
			notice := &api.EmojiLikeNotice{}
			if json.Unmarshal(*msgStr, notice) == nil {
				if n := parseEmojiLikeNotice(notice); n != nil {
					select {
					case r.emojiLike <- n:
					default:
					}
				}
			}
		case "group_upload":
			notice := &api.GroupUploadNotice{}
			if json.Unmarshal(*msgStr, notice) == nil {
				if n := parseGroupUploadNotice(notice); n != nil {
					select {
					case r.file <- n:
					default:
					}
				}
			}
		case "group_recall":
			fallthrough
		case "friend_recall":
			notice := &api.RecallNotice{}
			if json.Unmarshal(*msgStr, notice) == nil {
				if n := parseRecallNotice(notice); n != nil {
					select {
					case r.recall <- n:
					default:
					}
				}
			}
		case "notify":
			if header.SubType == "poke" {
				notice := &api.PokeNotify{}
				if json.Unmarshal(*msgStr, notice) == nil {
					if n := parsePokeNotify(notice); n != nil {
						select {
						case r.poke <- n:
						default:
						}
					}
				}
			}
		}
	case "message":
		msgJson := &api.MessageJson{}
		if json.Unmarshal(*msgStr, msgJson) != nil {
			return
		}
		msg, media, file := parseMsgJson(msgJson)
		if msg != nil {
			select {
			case r.message <- msg:
			default:
			}
		}
		if media != nil {
			switch media.Kind {
			case "record":
				select {
				case r.record <- media:
				default:
				}
			case "video":
				select {
				case r.video <- media:
				default:
				}
			}
		}
		if file != nil {
			select {
			case r.file <- file:
			default:
			}
		}
	}
}

func parseMsgJson(raw *api.MessageJson) (*Message, *MediaMessage, *FileMessage) {
	if raw == nil {
		return nil, nil, nil
	}

	result := Message{
		MsgID:   MsgID(raw.MessageID),
		UserID:  UserID(raw.Sender.UserID),
		GroupID: GroupID(raw.GroupID),
		Name:    raw.Sender.Nickname,
		Time:    raw.Time,
		Raw:     raw.RawMessage,
	}

	var (
		file  *FileMessage
		media *MediaMessage
	)

	if raw.Sender.Card != "" {
		result.GroupCard = raw.Sender.Card
	}

	switch raw.MessageType {
	case "private":
		result.ChatType = Private
	case "group":
		result.ChatType = Group
	default:
		result.ChatType = OtherChat
	}

	switch raw.Sender.Role {
	case "owner":
		result.GroupRole = GroupOwner
	case "admin":
		result.GroupRole = GroupAdmin
	case "member":
		result.GroupRole = GroupMember
	default:
		result.GroupRole = NotAGroup
	}

	for _, msg := range raw.Message {
		var jsonData map[string]any
		if err := json.Unmarshal(msg.Data, &jsonData); err != nil {
			return nil, nil, nil
		}

		switch msg.Type {
		case "reply":
			switch v := jsonData["id"].(type) {
			case string: // string
				replyId, _ := strconv.ParseUint(v, 10, 64)
				result.ReplyID = MsgID(replyId)
			case float64: // number
				result.ReplyID = MsgID(v)
			}
		case "text":
			if text, ok := jsonData["text"].(string); ok {
				result.Array = append(result.Array, TextItem(text))
			}
		case "at":
			var item AtItem
			switch v := jsonData["qq"].(type) {
			case string:
				if v == "all" {
					item = AtItem(AtAll)
				} else {
					qq, err := strconv.ParseInt(v, 10, 64)
					if err != nil {
						continue
					}
					item = AtItem(qq)
				}
			case float64:
				item = AtItem(v)
			}
			result.Array = append(result.Array, item)
		case "face":
			var item FaceItem
			switch v := jsonData["id"].(type) {
			case string:
				id, err := strconv.ParseInt(v, 10, 16)
				if err != nil {
					continue
				}
				item = FaceItem(id)
			case float64:
				item = FaceItem(v)
			}
			result.Array = append(result.Array, item)
		case "image":
			if url, ok := jsonData["url"].(string); ok {
				result.Array = append(result.Array, &ImageItem{
					Url: url,
				})
			}
		case "record", "video":
			media = &MediaMessage{Kind: msg.Type}
			if name, ok := jsonData["file"].(string); ok {
				media.FileName = name
			}
			if path, ok := jsonData["path"].(string); ok {
				media.Path = path
			}
			if url, ok := jsonData["url"].(string); ok {
				media.Url = url
			}
			switch v := jsonData["file_size"].(type) {
			case string:
				size, err := strconv.ParseInt(v, 10, 64)
				if err == nil {
					media.FileSize = size
				}
			case float64:
				media.FileSize = int64(v)
			}
		case "file":
			file = &FileMessage{Kind: "message"}
			if name, ok := jsonData["file"].(string); ok {
				file.FileName = name
			}
			if id, ok := jsonData["file_id"].(string); ok {
				file.FileID = id
			}
			switch v := jsonData["file_size"].(type) {
			case string:
				size, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					continue
				}
				file.FileSize = size
			case float64:
				file.FileSize = int64(v)
			}

		// case "record":
		// 	if path, ok := jsonData["path"].(string); ok {
		// 		result.Array = append(result.Array, &recordItem{
		// 			Path: path,
		// 		})
		// 	}
		// case "file":
		// 	result.Array = append(result.Array, &fileItem{
		// 		Data: string(msg.Data),
		// 	})
		// case "forward":
		// 	result.Array = append(result.Array, &forwardItem{
		// 		Data: string(msg.Data),
		// 	})
		// case "json":
		// 	result.Array = append(result.Array, &jsonItem{
		// 		Data: string(msg.Data),
		// 	})
		default:
			return nil, nil, nil
		}
	}
	if len(result.Array) == 0 && result.ReplyID == InvalidMsgID {
		result = Message{}
	}
	if file != nil {
		file.ChatType = result.ChatType
		file.MsgID = result.MsgID
		file.UserID = result.UserID
		file.Name = result.Name
		file.Time = result.Time
		file.GroupID = result.GroupID
		file.GroupCard = result.GroupCard
		file.GroupRole = result.GroupRole
	}
	if media != nil {
		media.ChatType = result.ChatType
		media.MsgID = result.MsgID
		media.UserID = result.UserID
		media.Name = result.Name
		media.Time = result.Time
		media.GroupID = result.GroupID
		media.GroupCard = result.GroupCard
		media.GroupRole = result.GroupRole
		media.Raw = result.Raw
	}
	sendMessage := len(result.Array) > 0 || result.ReplyID != InvalidMsgID
	if !sendMessage {
		return nil, media, file
	}
	return &result, media, file
}

func parseEmojiLikeNotice(raw *api.EmojiLikeNotice) *EmojiReaction {
	if raw == nil || len(raw.Likes) == 0 {
		return nil
	}

	notice := &EmojiReaction{
		GroupID:   GroupID(raw.GroupID),
		UserID:    UserID(raw.UserID),
		MessageID: MsgID(raw.MessageID),
		IsAdd:     raw.IsAdd,
		Count:     raw.Likes[0].Count,
	}

	id, err := strconv.ParseUint(raw.Likes[0].EmojiID, 10, 16)
	if err != nil {
		return nil
	}
	notice.FaceID = FaceID(id)

	if id < 1000 {
		notice.IsQFace = true
	} else {
		notice.IsQFace = false
		notice.EmojiRune = rune(id)
	}

	return notice
}

func parseRecallNotice(raw *api.RecallNotice) *RecallNotice {
	return &RecallNotice{
		ChatType:   Group,
		GroupID:    GroupID(raw.GroupID),
		UserID:     UserID(raw.UserID),
		OperatorID: UserID(raw.OperatorID),
		MessageID:  MsgID(raw.MessageID),
		Time:       raw.Time,
	}
}

func parsePokeNotify(raw *api.PokeNotify) *PokeNotify {
	if raw == nil {
		return nil
	}
	notify := &PokeNotify{
		ChatType: Group,
		GroupID:  GroupID(raw.GroupID),
		SenderID: UserID(raw.UserID),
		TargetID: UserID(raw.TargetID),
	}

	if notify.GroupID == InvalidGroup {
		notify.ChatType = Private
	}

	if list, ok := raw.RawInfo.([]any); ok {
		var txts []string
		for _, item := range list {
			if m, ok := item.(map[string]any); ok {
				if v, ok := m["txt"]; ok {
					if s, ok := v.(string); ok {
						txts = append(txts, s)
					}
				}
			}
		}
		if len(txts) > 0 {
			notify.Action = txts[0]
		}
		if len(txts) > 1 {
			notify.Suffix = txts[1]
		}
	}
	return notify
}

func parseGroupUploadNotice(raw *api.GroupUploadNotice) *FileMessage {
	if raw == nil {
		return nil
	}
	return &FileMessage{
		Kind:     "upload",
		ChatType: Group,
		GroupID:  GroupID(raw.GroupID),
		UserID:   UserID(raw.UserID),
		Time:     raw.Time,
		FileID:   raw.File.ID,
		FileName: raw.File.Name,
		FileSize: raw.File.Size,
		BusID:    raw.File.BusID,
	}
}
