package api

func SetGroupBan(c Client, groupID uint64, userID uint64, duration int) error {
	params := map[string]any{
		"group_id": groupID,
		"user_id":  userID,
		"duration": duration,
	}
	_, err := c.Send("set_group_ban", params)
	return err
}
