package api

import "encoding/json"

func GetFriendList(c Client) ([]FriendInfo, error) {
	data, err := c.Send("get_friend_list", nil)
	if err != nil {
		return nil, err
	}
	var resp []FriendInfo
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
