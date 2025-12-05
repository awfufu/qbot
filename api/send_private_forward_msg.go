package api

import "encoding/json"

func SendPrivateForwardMsg(c Client, userID uint64, messages []Segment, news []News, prompt, summary, source string) (int32, string, error) {
	params := map[string]any{
		"user_id":  userID,
		"messages": messages,
		"news":     news,
		"prompt":   prompt,
		"summary":  summary,
		"source":   source,
	}
	data, err := c.SendParams("send_private_forward_msg", params)
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
