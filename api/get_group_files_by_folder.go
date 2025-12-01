package api

import "encoding/json"

func GetGroupFilesByFolder(c Client, groupID uint64, folderID string) (*struct {
	Files   []GroupFile   `json:"files"`
	Folders []GroupFolder `json:"folders"`
}, error) {
	params := map[string]any{
		"group_id":  groupID,
		"folder_id": folderID,
	}
	data, err := c.Send("get_group_files_by_folder", params)
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
