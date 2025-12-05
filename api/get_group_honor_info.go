package api

import "encoding/json"

func GetGroupHonorInfo(c Client, groupID uint64, typeStr string) (*GroupHonorInfo, error) {
	params := map[string]any{
		"group_id": groupID,
		"type":     typeStr,
	}
	data, err := c.SendParams("get_group_honor_info", params)
	if err != nil {
		return nil, err
	}
	var resp GroupHonorInfo
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
