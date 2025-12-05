package api

import "encoding/json"

func GetGroupInfo(c Client, groupID uint64, noCache bool) (*GroupInfo, error) {
	params := map[string]any{
		"group_id": groupID,
		"no_cache": noCache,
	}
	data, err := c.SendParams("get_group_info", params)
	if err != nil {
		return nil, err
	}
	var resp GroupInfo
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
