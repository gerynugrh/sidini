package message

type Message struct {
	ID       uint64 `json:"ID"`
	Body     string `json:"body"`
	ParentID uint64 `json:"parentID"`
}
