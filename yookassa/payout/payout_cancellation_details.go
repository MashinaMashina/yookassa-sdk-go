package yoopayout

type PayoutCancellationDetails struct {
	Party  string `json:"party,omitempty"`
	Reason string `json:"reason,omitempty"`
}
