package response

type ListFriendsRecviceUpdateRespone struct {
	success    bool     `json:"success"`
	recipients []string `json:"recipients"`
}
