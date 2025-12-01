package api

import "encoding/json"

func GetUnidirectionalFriendList(c Client) ([]UnidirectionalFriendInfo, error) {
	data, err := c.Send("get_unidirectional_friend_list", nil)
	if err != nil {
		return nil, err
	}
	var resp []UnidirectionalFriendInfo
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
