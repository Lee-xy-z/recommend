package model

type Information struct {
	Action   string   `json:"action"`
	Info     Info     `json:"info"`
	Consumer []string `json:"consumer"`
}
