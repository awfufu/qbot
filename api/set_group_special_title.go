package api

func SetGroupSpecialTitle(c Client, groupID uint64, userID uint64, specialTitle string) error {
	params := map[string]any{
		"group_id":      groupID,
		"user_id":       userID,
		"special_title": specialTitle,
	}
	_, err := c.SendParams("set_group_special_title", params)
	return err
}
