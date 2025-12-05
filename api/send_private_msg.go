package api

import "encoding/json"

func SendPrivateMsg(c Client, userID uint64, message []Segment, autoEscape bool) (uint64, error) {
	params := map[string]any{
		"user_id":     userID,
		"message":     message,
		"auto_escape": autoEscape,
	}
	data, err := c.SendParams("send_private_msg", params)
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
