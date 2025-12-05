package api

func SetGroupWholeBan(c Client, groupID uint64, enable bool) error {
	params := map[string]any{
		"group_id": groupID,
		"enable":   enable,
	}
	_, err := c.SendParams("set_group_whole_ban", params)
	return err
}
