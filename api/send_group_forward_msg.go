package api

import "encoding/json"

func SendGroupForwardMsg(c Client, groupID uint64, messages []Segment, news []News, prompt, summary, source string) (int32, string, error) {
	params := map[string]any{
		"group_id": groupID,
		"messages": messages,
		"news":     news,
		"prompt":   prompt,
		"summary":  summary,
		"source":   source,
	}
	data, err := c.SendParams("send_group_forward_msg", params)
	if err != nil {
		return 0, "", err
	}
	var resp struct {
		MessageId int32  `json:"message_id"`
		ForwardId string `json:"forward_id"`
		ResId     string `json:"res_id"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return 0, "", err
	}
	if resp.ForwardId == "" {
		resp.ForwardId = resp.ResId
	}
	return resp.MessageId, resp.ForwardId, nil
}
