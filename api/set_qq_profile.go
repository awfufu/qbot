package api

func SetQqProfile(c Client, nickname, company, email, college, personalNote string) error {
	params := map[string]any{
		"nickname":      nickname,
		"company":       company,
		"email":         email,
		"college":       college,
		"personal_note": personalNote,
	}
	_, err := c.Send("set_qq_profile", params)
	return err
}
