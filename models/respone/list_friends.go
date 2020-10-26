package response

type ListFriendsRespone struct {
	success bool     `json:"success"`
	friends []string `json:"friends"`
	count   int      `json:"count"`
}
