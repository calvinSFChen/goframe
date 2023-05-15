package utils

import (
	"goframe/internal/consts"
	"time"
)

func TimeFormat(t string) (s string) {
	time, err := time.Parse(t, consts.TimeFormat)
	if err != nil {
		return
	}
	return time.String()
}
