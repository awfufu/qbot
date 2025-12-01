package api

func SetGroupPortrait(c Client, groupID uint64, file string, cache int) error {
	params := map[string]any{
		"group_id": groupID,
		"file":     file,
		"cache":    cache,
	}
	_, err := c.Send("set_group_portrait", params)
	return err
}
