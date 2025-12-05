package api

func CleanCache(c Client) error {
	_, err := c.SendParams("clean_cache", nil)
	return err
}
