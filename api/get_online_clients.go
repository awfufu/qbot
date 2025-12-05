package api

import "encoding/json"

func GetOnlineClients(c Client, noCache bool) ([]Device, error) {
	params := map[string]any{
		"no_cache": noCache,
	}
	data, err := c.SendParams("get_online_clients", params)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Clients []Device `json:"clients"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return resp.Clients, nil
}
