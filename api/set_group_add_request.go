package api

func SetGroupAddRequest(c Client, flag, subType string, approve bool, reason string) error {
	params := map[string]any{
		"flag":     flag,
		"sub_type": subType,
		"approve":  approve,
		"reason":   reason,
	}
	_, err := c.Send("set_group_add_request", params)
	return err
}
