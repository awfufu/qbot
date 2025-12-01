package api

func CleanCache(c Client) error {
	_, err := c.Send("clean_cache", nil)
	return err
}
