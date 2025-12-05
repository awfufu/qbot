package api

func SendGroupSign(c Client, groupID uint64) error {
	params := map[string]any{
		"group_id": groupID,
	}
	_, err := c.SendParams("send_group_sign", params)
	return err
}
