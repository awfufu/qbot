package api

func SetFriendAddRequest(c Client, flag string, approve bool, remark string) error {
	params := map[string]any{
		"flag":    flag,
		"approve": approve,
		"remark":  remark,
	}
	_, err := c.SendParams("set_friend_add_request", params)
	return err
}
