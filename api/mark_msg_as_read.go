package api

func MarkMsgAsRead(c Client, messageID int32) error {
	params := map[string]any{
		"message_id": messageID,
	}
	_, err := c.SendParams("mark_msg_as_read", params)
	return err
}
