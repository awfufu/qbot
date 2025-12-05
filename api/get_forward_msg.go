package api

import "encoding/json"

func GetForwardMsg(c Client, messageID string) ([]ForwardMsg, error) {
	params := map[string]any{
		"message_id": messageID,
	}
	data, err := c.SendParams("get_forward_msg", params)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Messages []ForwardMsg `json:"messages"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return resp.Messages, nil
}
