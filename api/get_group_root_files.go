package api

import "encoding/json"

func GetGroupRootFiles(c Client, groupID uint64) (*struct {
	Files   []GroupFile   `json:"files"`
	Folders []GroupFolder `json:"folders"`
}, error) {
	params := map[string]any{
		"group_id": groupID,
	}
	data, err := c.SendParams("get_group_root_files", params)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Files   []GroupFile   `json:"files"`
		Folders []GroupFolder `json:"folders"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
