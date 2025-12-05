package api

func FriendPoke(c Client, userID uint64) error {
	params := map[string]any{
		"user_id": userID,
	}
	_, err := c.SendParams("friend_poke", params)
	return err
}
