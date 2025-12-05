package api

import "encoding/json"

func GetCredentials(c Client, domain string) (*Credentials, error) {
	params := map[string]any{
		"domain": domain,
	}
	data, err := c.SendParams("get_credentials", params)
	if err != nil {
		return nil, err
	}
	var resp Credentials
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
