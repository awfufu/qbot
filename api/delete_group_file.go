package api

func DeleteGroupFile(c Client, groupID uint64, fileID string, busid int32) error {
	params := map[string]any{
		"group_id": groupID,
		"file_id":  fileID,
		"busid":    busid,
	}
	_, err := c.Send("delete_group_file", params)
	return err
}
