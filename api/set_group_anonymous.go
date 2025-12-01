package api

func SetGroupAnonymous(c Client, groupID uint64, enable bool) error {
	params := map[string]any{
		"group_id": groupID,
		"enable":   enable,
	}
	_, err := c.Send("set_group_anonymous", params)
	return err
}
