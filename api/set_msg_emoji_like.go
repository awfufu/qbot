package api

func SetMsgEmojiLike(c Client, messageID uint64, emojiID uint64, set bool) error {
	params := map[string]any{
		"message_id": messageID,
		"emoji_id":   emojiID,
		"set":        set,
	}
	_, err := c.SendParams("set_msg_emoji_like", params)
	return err
}
