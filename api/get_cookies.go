package api

import "encoding/json"

func GetCookies(c Client, domain string) (string, error) {
	params := map[string]any{
		"domain": domain,
	}
	data, err := c.Send("get_cookies", params)
	if err != nil {
		return "", err
	}
	var resp struct {
		Cookies string `json:"cookies"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return "", err
	}
	return resp.Cookies, nil
}
