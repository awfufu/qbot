package api

import "encoding/json"

func GetGroupFileSystemInfo(c Client, groupID uint64) (*GroupFileSystemInfo, error) {
	params := map[string]any{
		"group_id": groupID,
	}
	data, err := c.Send("get_group_file_system_info", params)
	if err != nil {
		return nil, err
	}
	var resp GroupFileSystemInfo
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
