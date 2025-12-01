package api

func SetGroupName(c Client, groupID uint64, groupName string) error {
	params := map[string]any{
		"group_id":   groupID,
		"group_name": groupName,
	}
	_, err := c.Send("set_group_name", params)
	return err
}
