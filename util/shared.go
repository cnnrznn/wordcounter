package util

type WCResponse struct {
	Wordcount []Count `json:"wordcount"`
	Time      float64 `json:"time"`
	URL       string  `json:"url"`
}

type Count struct {
	Key string `json:"key"`
	Val int    `json:"val"`
}
