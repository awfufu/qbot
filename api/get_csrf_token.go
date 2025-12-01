package api

import "encoding/json"

func GetCsrfToken(c Client) (int32, error) {
	data, err := c.Send("get_csrf_token", nil)
	if err != nil {
		return 0, err
	}
	var resp struct {
		Token int32 `json:"token"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return 0, err
	}
	return resp.Token, nil
}
