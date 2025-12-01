package api

import "encoding/json"

func GetVersionInfo(c Client) (*VersionInfo, error) {
	data, err := c.Send("get_version_info", nil)
	if err != nil {
		return nil, err
	}
	var resp VersionInfo
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
