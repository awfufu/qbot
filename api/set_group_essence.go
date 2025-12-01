package api

func SetGroupEssence(c Client, msgID uint64) error {
	params := map[string]any{
		"message_id": msgID,
	}
	_, err := c.Send("set_essence_msg", params)
	return err
}
