package api

import "encoding/json"

func OcrImage(c Client, imageID string) (*OcrResult, error) {
	params := map[string]any{
		"image": imageID,
	}
	data, err := c.SendParams("ocr_image", params)
	if err != nil {
		return nil, err
	}
	var resp OcrResult
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
