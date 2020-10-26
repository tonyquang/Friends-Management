package request

type ListFriendsReciveUpdateRequest struct {
	sender string `json:"sender"`
	text   string `json:"text"`
}
