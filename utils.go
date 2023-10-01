package simplequeue

import "encoding/hex"

func randomHex(n int) string {
	bytes := make([]byte, n)

	return hex.EncodeToString(bytes)
}
