package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var DisplayableMime = []string{
	"text/plain",
	"text/html",
	"text/css",
	"application/json",
	"application/javascript",
	"image/jpeg",
	"image/png",
	"image/gif",
	"audio/mpeg",
	"audio/ogg",
	"video/mp4",
	"video/ogg",
	"application/pdf",
}

func Pointer[T any](val T) *T { return &val }

func PointerArrayToValueArray[T any](arr []*T) []T {
	out := make([]T, len(arr))
	for i, a := range arr {
		out[i] = *a
	}
	return out
}

func ValueArrayToPointerArray[T any](arr []T) []*T {
	out := make([]*T, len(arr))
	for i, a := range arr {
		out[i] = Pointer(a)
	}
	return out
}

func SafeDereference[T any](val *T) T {
	if val == nil {
		return *new(T)
	}
	return *val
}

func Contains(arr []string, str string, like bool) bool {
	for _, a := range arr {
		if like {
			if strings.Contains(strings.ToLower(str), strings.ToLower(a)) {
				return true
			}
		} else {
			if a == str {
				return true
			}
		}
	}
	return false
}

func JSONMarshal(val interface{}) []byte {
	if val == nil {
		return []byte("{}")
	}
	v, _ := json.Marshal(val)
	return v
}

func JSONMarshalIndent(val interface{}) []byte {
	if val == nil {
		return []byte("{}")
	}
	v, _ := json.MarshalIndent(val, "", "  ")
	return v
}

func JSONUnMarshal[T any](data []byte) *T {
	if len(data) == 0 {
		return new(T)
	}
	var out T
	err := json.Unmarshal(data, &out)
	if err != nil {
		logrus.Error(fmt.Errorf("Error unmarshalling data: %s", err))
		return new(T)
	}
	return &out
}

func JSONArrayUnMarshal[T any](data []byte) []*T {
	if len(data) == 0 {
		return []*T{}
	}
	var out []*T
	err := json.Unmarshal(data, &out)
	if err != nil {
		logrus.Error(fmt.Errorf("Error unmarshalling data: %s", err))
		return []*T{}
	}
	return out
}

func GetENV(key string, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
