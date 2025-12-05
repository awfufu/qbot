package api

import "encoding/json"

func GetModelShow(c Client, model string) ([]ModelShow, error) {
	params := map[string]any{
		"model": model,
	}
	data, err := c.SendParams("_get_model_show", params)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Variants []ModelShow `json:"variants"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return resp.Variants, nil
}
