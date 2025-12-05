package api

import "encoding/json"

func SendGroupMsg(c Client, groupID uint64, message []Segment, autoEscape bool) (uint64, error) {

	params := map[string]any{
		"group_id":    groupID,
		"message":     message,
		"auto_escape": autoEscape,
	}
	data, err := c.SendParams("send_group_msg", params)
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
