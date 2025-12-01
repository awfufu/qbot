package api

func DeleteUnidirectionalFriend(c Client, userID uint64) error {
	params := map[string]any{
		"user_id": userID,
	}
	_, err := c.Send("delete_unidirectional_friend", params)
	return err
}
