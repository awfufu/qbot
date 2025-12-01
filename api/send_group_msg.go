package api

import "encoding/json"

func SendGroupMsg(c Client, groupID uint64, message string, autoEscape bool) (uint64, error) {
	if message == "" {
		message = " "
	}
	params := map[string]any{
		"group_id":    groupID,
		"message":     message,
		"auto_escape": autoEscape,
	}
	data, err := c.Send("send_group_msg", params)
	if err != nil {
		return 0, err
	}
	var resp struct {
		MessageId int32 `json:"message_id"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return 0, err
	}
	return uint64(resp.MessageId), nil
}
