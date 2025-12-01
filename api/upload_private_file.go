package api

func UploadPrivateFile(c Client, userID uint64, file, name string) error {
	params := map[string]any{
		"user_id": userID,
		"file":    file,
		"name":    name,
	}
	_, err := c.Send("upload_private_file", params)
	return err
}
