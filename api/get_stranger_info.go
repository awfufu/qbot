package api

import "encoding/json"

func GetStrangerInfo(c Client, userID uint64, noCache bool) (*StrangerInfo, error) {
	params := map[string]any{
		"user_id":  userID,
		"no_cache": noCache,
	}
	data, err := c.Send("get_stranger_info", params)
	if err != nil {
		return nil, err
	}
	var resp StrangerInfo
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
