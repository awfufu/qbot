package api

func DeleteGroupEssence(c Client, msgID uint64) error {
	params := map[string]any{
		"message_id": msgID,
	}
	_, err := c.Send("delete_essence_msg", params)
	return err
}
