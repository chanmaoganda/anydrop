package common

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func CheckSumSha256(filePath string) (string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return "", err
	}

	hash := sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	checksum := hex.EncodeToString(hash.Sum(nil))

	return checksum, nil
}
