package api

func DeleteGroupFileFolder(c Client, groupID uint64, folderID string) error {
	params := map[string]any{
		"group_id":  groupID,
		"folder_id": folderID,
	}
	_, err := c.SendParams("delete_group_file_folder", params)
	return err
}
