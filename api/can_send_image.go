package api

import "encoding/json"

func CanSendImage(c Client) (bool, error) {
	data, err := c.SendParams("can_send_image", nil)
	if err != nil {
		return false, err
	}
	var resp struct {
		Yes bool `json:"yes"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return false, err
	}
	return resp.Yes, nil
}
