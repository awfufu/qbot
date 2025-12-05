package api

import "encoding/json"

func GetGroupList(c Client, noCache bool) ([]GroupInfo, error) {
	params := map[string]any{
		"no_cache": noCache,
	}
	data, err := c.SendParams("get_group_list", params)
	if err != nil {
		return nil, err
	}
	var resp []GroupInfo
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
