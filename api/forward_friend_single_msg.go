package api

func ForwardFriendSingleMsg(c Client, userID uint64, messageID string) (uint64, error) {
	params := map[string]any{
		"user_id":    userID,
		"message_id": messageID,
	}
	_, err := c.SendParams("forward_friend_single_msg", params)
	if err != nil {
		return 0, err
	}
	return 0, nil
}
