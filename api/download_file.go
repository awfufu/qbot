package api

import "encoding/json"

func DownloadFile(c Client, url string, threadCount int, headers string) (string, error) {
	params := map[string]any{
		"url":          url,
		"thread_count": threadCount,
		"headers":      headers,
	}
	data, err := c.Send("download_file", params)
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
