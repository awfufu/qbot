package api

import "encoding/json"

func CanSendRecord(c Client) (bool, error) {
	data, err := c.Send("can_send_record", nil)
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
