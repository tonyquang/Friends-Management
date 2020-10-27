package response

type ListFriendsRespone struct {
	Success     bool     `json:"success"`
	ListFriends []string `json:"friends"`
	Count       int      `json:"count"`
}
