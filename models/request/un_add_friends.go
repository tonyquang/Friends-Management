package request

type AddFriendRequest struct {
	Friends []string `json:"friends"`
}

type UnFriendRequest struct {
	Friends []string `json:"unfriends"`
}
