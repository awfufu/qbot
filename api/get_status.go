package api

import "encoding/json"

func GetStatus(c Client) (*StatusInfo, error) {
	data, err := c.Send("get_status", nil)
	if err != nil {
		return nil, err
	}
	var resp StatusInfo
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
