package liveupdates

type Message struct {
	Op        string `json:"op"`
	Type      int32  `json:"type"`
	Condition string `json:"condition"`
}
