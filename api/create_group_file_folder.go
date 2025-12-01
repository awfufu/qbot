package api

func CreateGroupFileFolder(c Client, groupID uint64, name, parentID string) error {
	params := map[string]any{
		"group_id":  groupID,
		"name":      name,
		"parent_id": parentID,
	}
	_, err := c.Send("create_group_file_folder", params)
	return err
}
