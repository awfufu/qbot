package api

func SetGroupAnonymousBan(c Client, groupID uint64, anonymous, anonymousFlag string, duration int) error {
	params := map[string]any{
		"group_id":       groupID,
		"anonymous":      anonymous,
		"anonymous_flag": anonymousFlag,
		"duration":       duration,
	}
	_, err := c.SendParams("set_group_anonymous_ban", params)
	return err
}
