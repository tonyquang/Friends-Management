package response

type ListFriendsRecviceUpdateRespone struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}
