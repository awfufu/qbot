package api

func SetModelShow(c Client, model, modelShow string) error {
	params := map[string]any{
		"model":      model,
		"model_show": modelShow,
	}
	_, err := c.Send("_set_model_show", params)
	return err
}
