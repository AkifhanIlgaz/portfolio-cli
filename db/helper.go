package db

import "encoding/json"

func Serialize[T Transaction | Asset](data T) []byte {
	serialized, _ := json.Marshal(data)
	return serialized
}

func Deserialize[T Transaction | Asset](data []byte) T {
	var t T

	json.Unmarshal(data, &t)

	return t
}
