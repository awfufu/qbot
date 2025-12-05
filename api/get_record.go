package api

import "encoding/json"

func GetRecord(c Client, file, outFormat string) (string, error) {
	params := map[string]any{
		"file":       file,
		"out_format": outFormat,
	}
	data, err := c.SendParams("get_record", params)
	if err != nil {
		return "", err
	}
	var resp struct {
		File string `json:"file"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return "", err
	}
	return resp.File, nil
}
