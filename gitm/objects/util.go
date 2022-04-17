package objects

import (
	"crypto/sha1"
	"encoding/hex"
)

func Hash(bytes []byte) string {
	arr := sha1.Sum(bytes)
	return hex.EncodeToString(arr[:])
}
