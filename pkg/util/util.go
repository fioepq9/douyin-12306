package util

import (
	"bytes"
)

func Int64ToString(num int64) string {
	if num == 0 {
		return "0"
	}
	var (
		buf        bytes.Buffer
		numBuf     [19]byte
		i          = 19
		isNegative = num < 0
	)
	if isNegative {
		num = -num
		buf.WriteByte('-')
	}
	for num > 0 {
		i--
		next := num / 10
		numBuf[i] = byte('0' + num - next*10)
		num = next
	}
	buf.Grow(19 - i)
	for ; i < 19; i++ {
		buf.WriteByte(numBuf[i])
	}
	return buf.String()
}
