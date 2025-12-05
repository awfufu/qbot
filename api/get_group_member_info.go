package api

import "encoding/json"

func GetGroupMemberInfo(c Client, groupID uint64, userID uint64, noCache bool) (*GroupMemberInfo, error) {
	params := map[string]any{
		"group_id": groupID,
		"user_id":  userID,
		"no_cache": noCache,
	}
	data, err := c.SendParams("get_group_member_info", params)
	if err != nil {
		return nil, err
	}
	var resp GroupMemberInfo
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
