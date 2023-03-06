package common

import (
	"crypto/sha1"
	"encoding/base64"
)

// Generate ID by SHA-1 algorithm
func GenID(str string) string {
	hasher := sha1.New()
	bv := []byte(str)
	hasher.Write(bv)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
