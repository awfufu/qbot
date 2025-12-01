package api

import "encoding/json"

func SendMsg(c Client, messageType string, userID uint64, groupID uint64, message string, autoEscape bool) (int32, error) {
	params := map[string]any{
		"message_type": messageType,
		"user_id":      userID,
		"group_id":     groupID,
		"message":      message,
		"auto_escape":  autoEscape,
	}
	data, err := c.Send("send_msg", params)
	if err != nil {
		return 0, err
	}
	var resp struct {
		MessageId int32 `json:"message_id"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return 0, err
	}
	return resp.MessageId, nil
}
