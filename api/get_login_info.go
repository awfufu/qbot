package api

import "encoding/json"

func GetLoginInfo(c Client) (*LoginInfo, error) {
	data, err := c.Send("get_login_info", nil)
	if err != nil {
		return nil, err
	}
	var resp LoginInfo
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
