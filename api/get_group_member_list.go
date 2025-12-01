package api

import "encoding/json"

func GetGroupMemberList(c Client, groupID uint64, noCache bool) ([]GroupMemberInfo, error) {
	params := map[string]any{
		"group_id": groupID,
		"no_cache": noCache,
	}
	data, err := c.Send("get_group_member_list", params)
	if err != nil {
		return nil, err
	}
	var resp []GroupMemberInfo
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
