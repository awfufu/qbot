package api

import "encoding/json"

func GetGroupFileUrl(c Client, groupID uint64, fileID string, busid int32) (string, error) {
	params := map[string]any{
		"group_id": groupID,
		"file_id":  fileID,
		"busid":    busid,
	}
	data, err := c.Send("get_group_file_url", params)
	if err != nil {
		return "", err
	}
	var resp struct {
		Url string `json:"url"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return "", err
	}
	return resp.Url, nil
}
