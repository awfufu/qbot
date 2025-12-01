package api

func ReloadEventFilter(c Client, file string) error {
	params := map[string]any{
		"file": file,
	}
	_, err := c.Send("reload_event_filter", params)
	return err
}
