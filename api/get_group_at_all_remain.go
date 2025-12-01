package api

import "encoding/json"

func GetGroupAtAllRemain(c Client, groupID uint64) (bool, int32, int32, error) {
	params := map[string]any{
		"group_id": groupID,
	}
	data, err := c.Send("get_group_at_all_remain", params)
	if err != nil {
		return false, 0, 0, err
	}
	var resp struct {
		CanAtAll                 bool  `json:"can_at_all"`
		RemainAtAllCountForGroup int32 `json:"remain_at_all_count_for_group"`
		RemainAtAllCountForUin   int32 `json:"remain_at_all_count_for_uin"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return false, 0, 0, err
	}
	return resp.CanAtAll, resp.RemainAtAllCountForGroup, resp.RemainAtAllCountForUin, nil
}
