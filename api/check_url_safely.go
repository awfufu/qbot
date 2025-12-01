package api

import "encoding/json"

func CheckUrlSafely(c Client, url string) (int32, error) {
	params := map[string]any{
		"url": url,
	}
	data, err := c.Send("check_url_safely", params)
	if err != nil {
		return 0, err
	}
	var resp struct {
		Level int32 `json:"level"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return 0, err
	}
	return resp.Level, nil
}
