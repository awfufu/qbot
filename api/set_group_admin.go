package api

func SetGroupAdmin(c Client, groupID uint64, userID uint64, enable bool) error {
	params := map[string]any{
		"group_id": groupID,
		"user_id":  userID,
		"enable":   enable,
	}
	_, err := c.SendParams("set_group_admin", params)
	return err
}
