package api

import "encoding/json"

func GetEssenceMsgList(c Client, groupID uint64) ([]EssenceMsg, error) {
	params := map[string]any{
		"group_id": groupID,
	}
	data, err := c.Send("get_essence_msg_list", params)
	if err != nil {
		return nil, err
	}
	var resp []EssenceMsg
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
