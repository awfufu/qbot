package api

func GroupPoke(c Client, groupID uint64, userID uint64) error {
	params := map[string]any{
		"group_id": groupID,
		"user_id":  userID,
	}
	_, err := c.SendParams("group_poke", params)
	return err
}
