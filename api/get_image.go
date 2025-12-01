package api

import "encoding/json"

func GetImage(c Client, file string) (*ImageInfo, error) {
	params := map[string]any{
		"file": file,
	}
	data, err := c.Send("get_image", params)
	if err != nil {
		return nil, err
	}
	var resp ImageInfo
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
