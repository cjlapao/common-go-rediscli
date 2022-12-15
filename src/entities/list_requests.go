package entities

type RedisListSetRequest struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type RedisListCountResponse struct {
	Key   string `json:"key"`
	Count int64  `json:"count"`
}

type RedisListTrimRequest struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}
