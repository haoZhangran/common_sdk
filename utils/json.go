package utils

import (
	"github.com/bytedance/sonic"
)

func JsonMarshal(v any) string {
	bytes, err := sonic.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}
