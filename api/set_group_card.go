package api

func SetGroupCard(c Client, groupID uint64, userID uint64, card string) error {
	params := map[string]any{
		"group_id": groupID,
		"user_id":  userID,
		"card":     card,
	}
	_, err := c.SendParams("set_group_card", params)
	return err
}
