package request

type HandleUpdateRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}
