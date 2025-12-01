package api

func SetGroupKick(c Client, groupID uint64, userID uint64, rejectAddRequest bool) error {
	params := map[string]any{
		"group_id":           groupID,
		"user_id":            userID,
		"reject_add_request": rejectAddRequest,
	}
	_, err := c.Send("set_group_kick", params)
	return err
}
