package api

func UploadGroupFile(c Client, groupID uint64, file, name, folder string) error {
	params := map[string]any{
		"group_id": groupID,
		"file":     file,
		"name":     name,
		"folder":   folder,
	}
	_, err := c.SendParams("upload_group_file", params)
	return err
}
