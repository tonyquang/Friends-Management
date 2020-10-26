package request

type HandleUpdateRequest struct {
	requestor string `json:"requestor"`
	target    string `json:"target"`
}
