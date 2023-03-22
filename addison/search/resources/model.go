package resources

type Rep struct {
	Status string `json:"status"`
	Result Res    `json:"result"`
}

type Res struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
}