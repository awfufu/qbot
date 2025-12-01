package api

func SendGroupNotice(c Client, groupID uint64, content, image string) error {
	params := map[string]any{
		"group_id": groupID,
		"content":  content,
		"image":    image,
	}
	_, err := c.Send("_send_group_notice", params)
	return err
}
