package entities

type RedisSetKeyRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RedisGetKeyRequest struct {
	Key string `json:"key"`
}

type RedisGetKeyResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RedisGetKeysRequest struct {
	Pattern string `json:"pattern"`
}

type RedisGetKeysResponse struct {
	Keys []string `json:"keys"`
}
