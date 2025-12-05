package api

import "encoding/json"

func GetMsg(c Client, messageID int32) (*MessageJson, error) {
	params := map[string]any{
		"message_id": messageID,
	}
	data, err := c.SendParams("get_msg", params)
	if err != nil {
		return nil, err
	}
	var resp MessageJson
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
