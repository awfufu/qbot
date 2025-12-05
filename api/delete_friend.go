package api

func DeleteFriend(c Client, userID uint64) error {
	params := map[string]any{
		"user_id": userID,
	}
	_, err := c.SendParams("delete_friend", params)
	return err
}
