package common

import "os"

func Env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type Response interface {
	Code() int

	Headers() map[string]string

	Empty() bool
}
