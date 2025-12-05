package api

import "encoding/json"

func GetGroupMsgHistory(c Client, groupID uint64, messageSeq int32) ([]MessageJson, error) {
	params := map[string]any{
		"group_id":    groupID,
		"message_seq": messageSeq,
	}
	data, err := c.SendParams("get_group_msg_history", params)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Messages []MessageJson `json:"messages"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return resp.Messages, nil
}
