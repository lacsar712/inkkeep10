package pathsafe

import (
	"errors"
	"path/filepath"
	"strings"
)

func JoinUnder(root, rel string) (string, error) {
	if strings.TrimSpace(rel) == "" {
		return "", errors.New("empty path")
	}
	if filepath.IsAbs(rel) {
		return "", errors.New("absolute path")
	}
	return filepath.Join(root, rel), nil
}
