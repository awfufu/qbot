package api

func SetGroupLeave(c Client, groupID uint64, isDismiss bool) error {
	params := map[string]any{
		"group_id":   groupID,
		"is_dismiss": isDismiss,
	}
	_, err := c.SendParams("set_group_leave", params)
	return err
}
