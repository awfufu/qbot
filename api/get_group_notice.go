package api

import "encoding/json"

func GetGroupNotice(c Client, groupID uint64) ([]any, error) {
	params := map[string]any{
		"group_id": groupID,
	}
	data, err := c.SendParams("_get_group_notice", params)
	if err != nil {
		return nil, err
	}
	var resp []any
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
