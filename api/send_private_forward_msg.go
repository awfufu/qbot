package api

import "encoding/json"

func SendPrivateForwardMsg(c Client, userID uint64, messages []any) (int32, error) {
	params := map[string]any{
		"user_id":  userID,
		"messages": messages,
	}
	data, err := c.Send("send_private_forward_msg", params)
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
