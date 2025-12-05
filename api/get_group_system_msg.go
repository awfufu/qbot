package api

import "encoding/json"

func GetGroupSystemMsg(c Client) (*GroupSystemMsg, error) {
	data, err := c.SendParams("get_group_system_msg", nil)
	if err != nil {
		return nil, err
	}
	var resp GroupSystemMsg
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
